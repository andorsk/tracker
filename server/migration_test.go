package server

import (
	"fmt"
	"log"
	"testing"
	"tracker/config"
	gm "tracker/migration/db/gorp_migrations"
)

func TestCreateTable(t *testing.T) {
	conf, err := config.LoadConfig("../config/config.json")
	if err != nil {
		log.Fatal("Failed to load config")
	}
	fmt.Println("Initalizing with", conf.GetDB().User, ":", conf.GetDB().Password)
	s.Initialize(conf.GetDB().User, conf.GetDB().Password, "rest_api_example")

	gm.CreateTables(s.DB)

}
