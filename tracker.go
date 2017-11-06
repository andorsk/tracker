package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"tracker/config"
	"tracker/controller"
	"tracker/server"
)

var s server.Server

func Init() {
	s = server.Server{}
	conf, err := config.LoadConfig("./config/config.json")
	if err != nil {
		log.Fatal("Failed to load config")
	}

	fmt.Println("DB is ", conf.GetDB().DBName)
	s.Initialize(conf.GetDB().User, conf.GetDB().Password, conf.GetDB().DBName)
	mc := controller.MasterController{DB: s.DB, Router: s.Router}
	mc.InitializeRoutes(s.Router)

}

func Run() {
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func main() {
	Init()
	Run()
}
