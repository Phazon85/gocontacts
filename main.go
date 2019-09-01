package main

import "github.com/phazon85/multisql"

const (
	configFile = "dev.yaml"
	driverName = "postgres"
)

func main() {
	db := multisql.NewDBObject(configFile, driverName)
	
}
