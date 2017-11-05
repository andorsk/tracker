package heartbeat

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	hmi "tracker/model/heartbeat"
	"tracker/proto/heartbeat"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type HeartbeatController struct {
	Router *mux.Router
	DB     *sql.DB
}

//get the heartbeats for a specific user
func (h *HeartbeatController) GetHeartbeatsByUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Getting Heartbeats")
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["id"])
	fmt.Println("Looking for", userid)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User Id")
	}

	var heartbeats heartbeat.HeartbeatTrack

	_, err = hmi.Get(h.DB, userid)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}

		return
	}
	respondWithJSON(w, http.StatusOK, heartbeats)
}

func (h *HeartbeatController) PushHeartbeat(w http.ResponseWriter, r *http.Request) {

	var heartbeat heartbeat.Heartbeat
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&heartbeat); err != nil {
		log.Panic("Failed to decode heartbeat")
	}

	defer r.Body.Close()

	if err := hmi.Push(h.DB, heartbeat); err != nil {
		log.Panic("Failed to push heartbeat to database")
		respondWithError(w, http.StatusInternalServerError, "Failed to push heartbeat")
		return
	}
	respondWithJSON(w, http.StatusCreated, heartbeat)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (h *HeartbeatController) InitializeRoutes(r *mux.Router) {
	r.HandleFunc("/hb-api/{id:[0-9]+}", h.GetHeartbeatsByUser).Methods("GET")
	r.HandleFunc("/hb-api", h.PushHeartbeat).Methods("POST")
}
