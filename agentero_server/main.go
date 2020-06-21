//go:generate protoc -I ../agentero --go_out=plugins=grpc:../agentero ../agentero/agentero.proto
package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/rafmsou/agentero/agentero"
	"github.com/rafmsou/agentero/importer"
	"github.com/rafmsou/agentero/models"
	"github.com/rafmsou/agentero/services"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedAPIServer
}

func (s *server) GetContactAndPoliciesByID(
	ctx context.Context,
	in *pb.RequestById,
) (*pb.ContactPoliciesReply, error) {
	log.Printf("GetContactAndPoliciesById for userId: %v", in.GetUserId())

	user, err := services.GetUserByID(in.GetUserId())
	if err != nil {
		return nil, err
	}

	return getPoliciesForContact(user)
}

func (s *server) GetContactsAndPoliciesByMobileNumber(
	ctx context.Context,
	in *pb.RequestByMobileNumber,
) (*pb.ContactPoliciesReply, error) {
	log.Printf("GetContactAndPoliciesById for mobileNumber: %v", in.GetMobileNumber())

	user, err := services.GetUserByMobileNumber(in.GetMobileNumber())
	if err != nil {
		return nil, err
	}
	return getPoliciesForContact(user)
}

func getPoliciesForContact(user *models.User) (*pb.ContactPoliciesReply, error) {
	policies, err := services.GetPoliciesByMobileNumber(user.MobileNumber)
	if err != nil {
		return nil, err
	}

	pbPolicies := []*pb.Policy{}
	for _, p := range policies {
		pbPolicies = append(pbPolicies, &pb.Policy{
			Type:    p.Type,
			Premium: p.Premium,
		})
	}
	return &pb.ContactPoliciesReply{
		Name:         user.Name,
		MobileNumber: user.MobileNumber,
		Policies:     pbPolicies,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	amsServerAddress := flag.String("ams-api-url", "", "AMS Server Address")
	importSchedulePeriod := flag.Duration("schedule-period", 0, "Interval which to run data import")
	flag.Parse()

	if *amsServerAddress != "" {
		for _, agentID := range services.GetAgentsIDsToImport() {
			err = importer.Run(*amsServerAddress, agentID)
			if err != nil {
				log.Fatalf("failed to import data from %s for agentID %d", *amsServerAddress, agentID)
			}
			if *importSchedulePeriod > 0 {
				go importer.RunAtInterval(*amsServerAddress, agentID, *importSchedulePeriod)
			}
		}
	}

	s := grpc.NewServer()
	pb.RegisterAPIServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
