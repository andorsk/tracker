package heartbeat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"tracker/config"
	"tracker/proto/heartbeat"
	"tracker/server"
	//	logger "github.com/sirupsen/logrus"
)

var s server.Server
var hb HeartbeatController

const UUIDinstertion = `a427abe0-d359-4bbd-be70-e5a6b83defed`
const UUIDinstertion2 = `b427abe0-d359-4bbd-be70-e5a6b83defe7`

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

func TestAddHeartbeatToNonExistentUser(t *testing.T) {
	clearTable()

	jsonStr := fmt.Sprintf(`{"latitude": %v, "longitude": %v, "timestamp":%v, "userId": %v})`, 1, 1, 1, 45)
	payload := []byte(jsonStr)
	req, _ := http.NewRequest("POST", "/hb-api", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestAddHeartbeatToExistent(t *testing.T) {
	addUsers(1)
	jsonStr := fmt.Sprintf(`{"latitude": %v, "longitude": %v, "timestamp":%v, "userId": %v})`, 1, 1, 1, 1)
	payload := []byte(jsonStr)
	req, _ := http.NewRequest("POST", "/hb-api", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

}
func TestAddHeartbeatAgainWithAnalyticThere(t *testing.T) {
	jsonStr := fmt.Sprintf(`{"latitude": %v, "longitude": %v, "timestamp":%v, "userId": %v})`, 1, 3, 3, 1)
	payload := []byte(jsonStr)
	req, _ := http.NewRequest("POST", "/hb-api", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
}

func addHeartBeat() {
	statement := fmt.Sprintf("INSERT INTO tracks(Uuid, UserId, Starttime) VALUES ('%s', '%v', '%v')", UUIDinstertion, 1, 2)
	s.DB.Exec(statement)
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}
	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES('%s', %d)", ("User " + strconv.Itoa(i+1)), ((i + 1) * 10))
		s.DB.Exec(statement)
	}
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
