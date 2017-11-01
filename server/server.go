package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
		log.Fatal(err)
	}
	s.Router = mux.NewRouter()
	s.InitializeRoutes()
}

func (s *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.Router))
}

func (s *Server) AccessRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("You made it to the end"))
}

func (s *Server) InitializeRoutes() {
	fmt.Println("Initalizing Routes")
	s.Router.HandleFunc("/", s.AccessRoot)
}
