package main

import (
	"flag"
	"log"

	"github.com/souravsk/BookMyTicket/src/eventservices/rest"
	"github.com/souravsk/BookMyTicket/src/lib/configuration"
	"github.com/souravsk/BookMyTicket/src/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `.\app\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	log.Println("Connecting to database")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful... ")
	//RESTful API start
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
