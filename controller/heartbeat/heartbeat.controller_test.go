package heartbeat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"tracker/config"
	"tracker/proto/heartbeat"
	"tracker/server"
	//	logger "github.com/sirupsen/logrus"
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

	req, _ := http.NewRequest("GET", "/hb-api?UserId=45", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string

	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "User not found" {
		t.Errorf("Expected User not found. Got %s", m["error"])
	}

}

func TestAddHeartbeat(t *testing.T) {
	clearTable()

	payload := []byte(`{"Logitude": "12", "Latitude": 1, "UserId": 1, "Timestamp": 1}`)
	req, _ := http.NewRequest("POST", "/hb-api", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	fmt.Println(m)

}

const UUIDinstertion = `a427abe0-d359-4bbd-be70-e5a6b83defed`

func addHeartBeat() {

	statement := fmt.Sprintf("INSERT INTO tracks(Uuid, UserId, Starttime) VALUES ('%s', '%v', '%v')", UUIDinstertion, 1, 2)
	s.DB.Exec(statement)
}

func TestGetHeartbeatForCorrectUser(t *testing.T) {
	clearTable()
	addHeartBeat()
	req, _ := http.NewRequest("GET", "/hb-api?UserId=1", nil)
	response := executeRequest(req)

	err := checkResponseCode(t, http.StatusOK, response.Code)
	if err != nil {
		t.Errorf("Response: %v", response.Body.String())
	}

	var tracks []heartbeat.HeartbeatTrack
	if err := json.Unmarshal(response.Body.Bytes(), &tracks); err != nil {
		log.Panic("Failed to Unrmashal JSON", response.Body.String())
	}

	if len(tracks) > 1 {
		t.Error("Only supposed to be one entry. Continue")
	}
	for _, track := range tracks {
		if track.UserId != 1 {
			t.Error("Failed to reqeust correct user id")
		}
		if string(track.Uuid.Value) != string(UUIDinstertion) {
			t.Errorf("Failed to get correct uuid value \n %s : %s", track.Uuid.Value, UUIDinstertion)
		}
	}
}

func TestMalformedPush(t *testing.T) {
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) error {
	if expected != actual {
		t.Errorf("Expected was %d. Got %d", expected, actual)
		return errors.New("Falied Test")
	}
	return nil
}

func clearTable() {
	s.DB.Exec("DELETE FROM heartbeat")
	s.DB.Exec("ALTER TABLE heartbeat AUTO_INCREMENT = 1")
	s.DB.Exec("DELETE FROM tracks")
}
