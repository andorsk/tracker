package heartbeat

import (
	"database/sql"
	"fmt"
	"log"
	"tracker/proto/heartbeat"
	puuid "tracker/proto/uuid"
)

//Theres a dependency between a heartbeat and a heartbeat track. You cannot have a heartbeat point to a heartbeat tracka and the heartbeat track not exist.

func Push(db *sql.DB, hb heartbeat.Heartbeat) error {

	statement := fmt.Sprintf("INSERT INTO heartbeat (UserId, Location, Longitude, Latitude, Timestamp) VALUES ('%d', '%s', '%v', '%v', '%v')", hb.UserId, hb.Location, hb.Longitude, hb.Latitude, hb.Timestamp)

	_, err := db.Exec(statement)

	if err != nil {
		log.Panic("Query was: ", statement)
		log.Panic("Failed to push heartbeat to model:", err.Error())
		return nil
	}

	return nil
}

func Get(db *sql.DB, column, criteria string) ([]heartbeat.HeartbeatTrack, error) {
	statement := fmt.Sprintf("SELECT * FROM tracks WHERE %s = %s", column, string(criteria))

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	hbtracks := []heartbeat.HeartbeatTrack{}

	for rows.Next() {
		var hbt heartbeat.HeartbeatTrack
		var uid []uint8
		var hbs sql.NullString

		if err := rows.Scan(&uid, &hbs, &hbt.Starttime, &hbt.UserId); err != nil {
			return nil, err
		}

		var t puuid.UUID
		t.Value = uid
		hbt.Uuid = &t
		hbtracks = append(hbtracks, hbt)
	}

	if len(hbtracks) == 0 {
		return nil, sql.ErrNoRows
	}

	return hbtracks, nil
}

func DeleteHeartbeats(db *sql.DB, id int) error {
	return nil
}
