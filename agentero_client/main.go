package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rafmsou/agentero/agentero"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "agentero"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call GetContactAndPoliciesByID and print its result
	r, err := c.GetContactAndPoliciesByID(ctx, &pb.RequestById{UserId: 2})
	if err != nil {
		log.Fatalf("error calling GetContactAndPoliciesById: %v", err)
	}
	log.Printf("GetContactAndPoliciesByID: %s", r)

	// Call GetContactsAndPoliciesByMobileNumber and print its result
	r, err = c.GetContactsAndPoliciesByMobileNumber(
		ctx,
		&pb.RequestByMobileNumber{MobileNumber: "1234567890"},
	)
	if err != nil {
		log.Fatalf("error calling GetContactsAndPoliciesByMobileNumber: %v", err)
	}
	log.Printf("GetContactsAndPoliciesByMobileNumber: %s", r)
}
