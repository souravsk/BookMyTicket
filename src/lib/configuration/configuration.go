package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/souravsk/BookMyTicket/src/lib/persistence/dblayer"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault    = "localhost:8888"
)

// ServiceConfig is a struct that holds the configuration for the service.
type ServiceConfig struct {
	Databasetype    dblayer.DBTYPE `json:"databasetype"`
	DBConnection    string         `json:"dbconnection"`
	RestfulEndpoint string         `json:"restfulapi_endpoint"`
}

// ExtractConfiguration reads a configuration file and returns a ServiceConfig.
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		Databasetype:    DBTypeDefault,
		DBConnection:    DBConnectionDefault,
		RestfulEndpoint: RestfulEPDefault,
	}
	fmt.Println("Deafult configuration: ", conf)
	fmt.Println("Reading configuration from file: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.", err)
		return conf, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&conf)
	fmt.Println("Returning configuration: ", conf)
	return conf, err
}
