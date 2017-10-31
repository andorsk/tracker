package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Heartbeat struct {
	Longitude float64
	Latitude  float64
	Timestamp int64
}

var heartbeats []Heartbeat

func AppendHeartbeat(w http.ResponseWriter, r *http.Request) {
	heartbeats = append(heartbeats, Heartbeat{Longitude: 4, Latitude: 4, Timestamp: 4})
	b, _ := ioutil.ReadAll(r.Body)
	w.Write(b)
}

func GetHeartbeats(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeats)
	fmt.Println("Here you go")
}

func genFakeData() {
	heartbeats = append(heartbeats, Heartbeat{Longitude: 1, Latitude: 2, Timestamp: 3})
}

func main() {
	genFakeData()
	router := mux.NewRouter()
	router.HandleFunc("/get", GetHeartbeats).Methods("GET")
	router.HandleFunc("/append", AppendHeartbeat).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Serving Webpage")
	log.Fatal(http.ListenAndServe(":12345", router))
}
