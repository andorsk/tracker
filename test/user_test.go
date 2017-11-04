package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"tracker/config"
)

var s Server

func TestMain(m *testing.M) {
	s = Server{}
	conf, err := config.LoadConfig("../config/config.json")
	if err != nil {
		log.Fatal("Failed to load config")
	}
	fmt.Println("Initalizing with ", conf.GetDB().User, " : ", conf.GetDB().Password)
	s.Initialize(conf.GetDB().User, conf.GetDB().Password, "rest_api_example")
	ensureTableExists()

	code := m.Run()
	clearTable()
	os.Exit(code)

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

func TestCreateUser(t *testing.T) {
	clearTable()

	payload := []byte(`{"name": "test", "age": 30}`)

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	//run checks on create user
	if m["name"] != "test" {
		t.Errorf("Expected test user got %v", m["name"])
	}

	if m["age"] != 30.0 {
		t.Errorf("Expected age of 30 got %v", m["age"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected id to be 1 got %v", m["id"])
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

func ensureTableExists() {
	if _, err := s.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
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

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users( id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
)`
