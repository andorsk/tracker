package heartbeat

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type HeartbeatController struct {
	Router *mux.Router
	DB     *sql.DB
}

//get the heartbeats for a specific user
func (h *HeartbeatController) GetHeartbeats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["UserId"])
	check(err, "Failed to parse DB string")
	statement := fmt.Sprintf("SELECT * FROM heartbeats WHERE uid = %s", id)
	fmt.Println("Querying", statement)
}

func check(err error, message interface{}) {
	if err != nil {
		log.Panic("System error", err.Error(), message)
	}
}
func (h *HeartbeatController) PushHeartbeat(w http.ResponseWriter, r *http.Request) {

}

func (h *HeartbeatController) InitializeRoutes(r *mux.Router) {
	r.HandleFunc("/heartbeats", h.GetHeartbeats).Methods("GET")
	r.HandleFunc("/pushbeat", h.PushHeartbeat).Methods("POST")
}
