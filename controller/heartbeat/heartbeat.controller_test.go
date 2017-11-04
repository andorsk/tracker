package heartbeat

import (
	"fmt"
	"log"
	"testing"
	"tracker/proto/config"
)

var s Server

func TestMain(t *testing.T) {
	s.Initalize("mysql", "root", "c0raline")
	conf := config.Config.LoadConfig("../../config/config.json")
	fmt.Println(conf)
}

func TestGetHeartbeatForIncorrectUser() {
}

func TestGetHeartbeatForCorrectUser() {
}

func TestEmptyTable() {
}

func TestAddHeartbeat() {
}

func TestBadPush() {
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS heartbeats( id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp BIGINT NOT NULL,
    location BLOB NOT NULL,
    iud INT NOT NULL
)`

func ensureTableExists() {
	_, err := hb.DB.Exec(tableCreationQuery)
	if err != nil {
		log.Fatal("Table failed to load")
	}
}
