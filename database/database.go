package database

import (
	"log"

	"github.com/hashicorp/go-memdb"
)

var db *memdb.MemDB

// InitializeDatabase creates the DB schema and initializes a singleton instance of the database
func InitializeDatabase() error {
	log.Print("Initializing database")

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"user": &memdb.TableSchema{
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"mobile_number": &memdb.IndexSchema{
						Name:    "mobile_number",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MobileNumber"},
					},
					"agent_id": &memdb.IndexSchema{
						Name:    "agent_id",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "AgentID"},
					},
				},
			},
			"policy": &memdb.TableSchema{
				Name: "policy",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
					"premium": &memdb.IndexSchema{
						Name:    "premium",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Premium"},
					},
					"type": &memdb.IndexSchema{
						Name:    "type",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Type"},
					},
					"mobile_number": &memdb.IndexSchema{
						Name:    "mobile_number",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MobileNumber"},
					},
				},
			},
		},
	}

	// Create a new database
	newdb, err := memdb.NewMemDB(schema)
	if err != nil {
		return err
	}
	db = newdb

	return nil
}

// GetInstance returns a database instance
func GetInstance() (*memdb.MemDB, error) {
	if db == nil {
		err := InitializeDatabase()
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

// DestroyInstance clean current db instance
func DestroyInstance() {
	db = nil
}
