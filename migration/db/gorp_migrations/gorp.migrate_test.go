package gorp_migrations

import (
	"fmt"
	"log"
	"testing"
	"tracker/config"
	"tracker/server"
)

var s server.Server

func TestCreateTable(t *testing.T) {
	conf, err := config.LoadConfig("../../../config/config.json")
	if err != nil {
		log.Fatal("Failed to load config")
	}
	fmt.Println("Initalizing with", conf.GetDB().User, ":", conf.GetDB().Password)
	s.Initialize(conf.GetDB().User, conf.GetDB().Password, "rest_api_example")
	DropTables(s.DB)
	CreateTables(s.DB)
}
