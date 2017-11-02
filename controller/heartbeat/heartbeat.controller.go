package heartbeat

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

type HeartbeatController struct {
	Router *mux.Router
	DB     *sql.DB
}

func (h *HeartbeatController) GetHeartbeats(w mux.ResponseWriter, r *http.Request) {
	vars = mux.Vars(r)
	id := strconv.Atoi(vars["uid"])
	statement := fmt.Sprintf("SELECT * FROM heartbeats WHERE uid = %s", id)
	fmt.Println("Querying", statement)

}

func (h *HeartbeatController) PushHeartbeat(w mux.ResponseWriter, r *http.Request) {

}

func (h *HeartbeatController) InitializeRoutes(r *mux.Router) {
	r.HandleFunc("/heartbeats", h.GetHeartbeats()).Methods("GET")
	r.HandleFunc("/pushbeat", h.PushHeartbeat()).Methods("POST")
}
