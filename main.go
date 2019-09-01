package main

import "github.com/phazon85/multisql"

const (
	configFile = "dev.yaml"
	driverName = "postgres"
)

func main() {

	//load DB connection
	db := multisql.NewDBObject(configFile, driverName)
	
	router := handler.
}
