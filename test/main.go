package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var members = []Member{Member{"someuser", "someuser@somedomain.com"}}

type Member struct {
	Login string
	Email string
}

func getMembersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Getting")
	j, _ := json.Marshal(members)
	w.Write(j)
}

func postMembersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m Member
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &m)

	members = append(members, m)
	fmt.Println("This is a test")
	j, _ := json.Marshal(m)
	w.Write(j)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive":true}`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/members", getMembersHandler).Methods("GET")
	r.HandleFunc("/members", postMembersHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
