package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"tracker/controller/heartbeat"
	"tracker/controller/user"

	"github.com/gorilla/mux"
)

type MasterController struct {
	Router *mux.Router
	DB     *sql.DB
}

func (m *MasterController) AccessRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("You made it to the end"))
}

func (m *MasterController) InitializeRoutes(r *mux.Router) {
	fmt.Println("Initalizing Routes")
	m.Router.HandleFunc("/", m.AccessRoot)
	uc := user.UserController{Router: m.Router, DB: m.DB}
	uc.InitializeRoutes(m.Router)
	hb := heartbeat.HeartbeatController{Router: m.Router, DB: m.DB}
	hb.InitializeRoutes(m.Router)

}