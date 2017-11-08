package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"tracker/config"
	"tracker/server"

	log "github.com/sirupsen/logrus"
)

var s server.Server
var uc UserController

func init() {
	s = server.Server{}
	conf, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		log.Fatal("Failed to load config")
	}

	fmt.Println("DB is ", conf.GetDB().DBName)
	s.Initialize(conf.GetDB().User, conf.GetDB().Password, conf.GetDB().DBName)
	uc = UserController{DB: s.DB, Router: s.Router}
	uc.InitializeRoutes(s.Router)

}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}

}

func TestGetNonExistentUser(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/user/45", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string

	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

func TestGetUserById(t *testing.T) {
	clearTable()
	addUsers(2)
	req, _ := http.NewRequest("GET", "/user/2", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	fmt.Println(response.Body.String())
}

func TestCreateUser(t *testing.T) {
	clearTable()

	payload := []byte(`{"Name": "test user", "Age": 30}`)

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["Name"] != "test user" {
		t.Errorf("Expected test user got %v", m["Name"])
	}

	if m["Age"] != 30.0 {
		t.Errorf("Expected age of 30 got %v", m["Age"])
	}

	if m["UserId"] != 1.0 {
		t.Errorf("Expected id to be 1 got %v", m["UserId"])
	}

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

func addUsers(count int) {
	if count < 1 {
		count = 1
	}
	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES('%s', %d)", ("User " + strconv.Itoa(i+1)), ((i + 1) * 10))
		s.DB.Exec(statement)
	}
}

func clearTable() {
	s.DB.Exec("DELETE FROM users")
	s.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}
