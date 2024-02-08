package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/souravsk/BookMyTicket/src/eventservices/rest"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration file")
	flag.Parse()
	// Extract the configuration from the configuration file
	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("Connecting to database")
	// Create a new database handler
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	//RESTful API start using Gin
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))

}
