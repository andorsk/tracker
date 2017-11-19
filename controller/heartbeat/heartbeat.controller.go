package heartbeat

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	hmi "tracker/model/heartbeat"
	"tracker/model/uuid"
	phb "tracker/proto/heartbeat"

	"encoding/json"
	umi "tracker/model/user"
	puid "tracker/proto/uuid"

	logger "github.com/sirupsen/logrus"

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
	if len(vars["UserId"]) == 0 {
		logger.Error("Error in request", r.URL.Query)
		return
	}
	uid, err := strconv.Atoi(vars["UserId"][0])
	fmt.Println("Getting user id ", uid)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User Id. Please make sure to specify number for user id")
		return
	}

	if len(vars["UserId"]) > 1 {
		respondWithError(w, http.StatusBadRequest, "Please make sure to only have one user id")
		return
	}

	var heartbeats []phb.HeartbeatTrack

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

//TODO: Check if it's an invalid heartbeat.
func (h *HeartbeatController) PushHeartbeat(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)

	var hb phb.Heartbeat
	if err := dec.Decode(&hb); err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	if h.checkUserExists(w, hb.UserId) == false {
		logger.Warning("User does not exist. Failed to push heartbeat")
		respondWithError(w, http.StatusNotFound, "User Not Found")
		return
	}

	_, err := h.getLatestUUID(hb)
	if err != nil {
		logger.Warning("No track found. Creating tracks....")
		h.addHeartbeatTrack(hb)
	}

	if err := hmi.Push(h.DB, hb); err != nil {
		logger.Warning("Failed to push heartbeat to database")
		respondWithError(w, http.StatusInternalServerError, "Failed to push heartbeat")
		return
	}
	respondWithJSON(w, http.StatusCreated, hb)
}

func (h *HeartbeatController) checkUserExists(w http.ResponseWriter, id int64) bool {
	_, err := umi.GetByUserId(h.DB, id)
	if err != nil {
		return false
	}
	return true
}

func (h *HeartbeatController) addHeartbeatTrack(hb phb.Heartbeat) {

	uuidst, err := uuid.NewUUIDString()
	if err != nil {
		logger.Panic("Failed to Create UUID")
	}
	statement := fmt.Sprintf("INSERT INTO tracks(Uuid, Starttime, UserId) VALUES ('%v', '%v', '%v')", uuidst, hb.Timestamp, hb.UserId)
	h.DB.Exec(statement)
}

func AddHeartbeatTrack(db *sql.DB, hb phb.Heartbeat) {

	uuidst, err := uuid.NewUUIDString()
	if err != nil {
		logger.Panic("Failed to Create UUID")
	}
	statement := fmt.Sprintf("INSERT INTO tracks(Uuid, Starttime, UserId) VALUES ('%v', '%v', '%v')", uuidst, hb.Timestamp, hb.UserId)
	db.Exec(statement)
}

func (h *HeartbeatController) getLatestUUID(hb phb.Heartbeat) (*puid.UUID, error) {

	var heartbeats []phb.HeartbeatTrack
	ustr := strconv.FormatInt(hb.UserId, 16)

	heartbeats, err := hmi.Get(h.DB, "UserId", ustr)

	if err != nil {
		return &puid.UUID{}, err
	}
	return heartbeats[len(heartbeats)-1].Uuid, nil
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
	logger.Info("Registered Heartbeat Routes")
	s := r.PathPrefix("/hb-api").Subrouter()
	s.HandleFunc("", h.GetHeartbeatsByUser).Methods("GET")
	s.HandleFunc("", h.PushHeartbeat).Methods("POST")
}
