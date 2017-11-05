package heartbeat

import (
	"database/sql"
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

//get the heartbeats for a specific user. ONLY one user at a time.
func (h *HeartbeatController) GetHeartbeatsByUser(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()
	_, err := strconv.Atoi(vars["UserId"][0])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User Id. Please make sure to specify number for user id")
	}

	if len(vars["UserId"]) > 1 {
		respondWithError(w, http.StatusBadRequest, "Please make sure to only have one user id")
	}

	var heartbeats []heartbeat.HeartbeatTrack

	heartbeats, err = hmi.Get(h.DB, "UserId", vars["UserId"][0])

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
	s := r.PathPrefix("/hb-api").Subrouter()
	s.HandleFunc("", h.GetHeartbeatsByUser).Methods("GET")
	//s.HandleFunc("", h.PushHeartbeat).Methods("POST")
}
