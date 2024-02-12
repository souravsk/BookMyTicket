package dblayer

import (
	"fmt"

	"github.com/souravsk/BookMyTicket/src/lib/persistence"
	"github.com/souravsk/BookMyTicket/src/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DOCUMENTDB DBTYPE = "documentdb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	fmt.Printf("Options: %s\n", options)
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
