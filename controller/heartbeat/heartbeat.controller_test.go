package heartbeat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"tracker/config"
	"tracker/server"
)

var s server.Server
var hb HeartbeatController

func init() {
	s = server.Server{}
	conf, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		log.Fatal("Failed to load config")
	}

	fmt.Println("DB is ", conf.GetDB().DBName)
	s.Initialize(conf.GetDB().User, conf.GetDB().Password, conf.GetDB().DBName)
	hb = HeartbeatController{DB: s.DB, Router: s.Router}
	hb.InitializeRoutes(s.Router)

}

func TestGetHeartbeatForIncorrectUser(t *testing.T) {

	req, _ := http.NewRequest("GET", "/hb-api/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string

	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "User not found" {
		t.Errorf("Expected User not found. Got %s", m["error"])
	}

}

func TestGetHeartbeatForCorrectUser(t *testing.T) {
}

func TestEmptyTable(t *testing.T) {
}

func TestAddHeartbeat(t *testing.T) {
}

func TestBadPush(t *testing.T) {
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected was %d. Got %d", expected, actual)
	}
}

func clearTable() {
	s.DB.Exec("DELETE FROM heartbeat")
	s.DB.Exec("ALTER TABLE heartbeat AUTO_INCREMENT = 1")
}
