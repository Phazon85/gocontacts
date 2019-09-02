package main

import (
	"github.com/phazon85/go_contacts/handler"
	"github.com/phazon85/multisql"
)

const (
	configFile = "dev.yaml"
	driverName = "postgres"
)

func main() {

	//load DB connection
	db := multisql.NewDBObject(configFile, driverName)

	//intializes HTTP router
	handler := handler.NewEntryHandler(db)
}
