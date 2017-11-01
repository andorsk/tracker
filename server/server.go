package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"tracker/controller/user"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

func (s *Server) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s:%s", user, password, dbname)

	var err error
	s.DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		log.Fata(err)
	}
	s.Router = mux.NewRouter()
	s.InitializeRoutes()
}

func (s *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (s *Server) InitializeRoutes() {
	fmt.Println("Initalizing Routes")
	s.Router.HandleFunc("/user/*", user.UserController.InitializeRoutes(s.Router))
}
