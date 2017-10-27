package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Here you go")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", MainHandler).Methods("POST")
	fmt.Println("Serving Webpage")
	log.Fatal(http.ListenAndServe(":12345", router))
}
