package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

func (s *Server) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	s.DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}
	s.Router = mux.NewRouter()
	logger.Info("Successfully Opened up Connection to Database:", dbname)
}

func (s *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
