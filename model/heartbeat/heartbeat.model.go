package heartbeat

import (
	"database/sql"
	"fmt"
	"log"
	"tracker/proto/heartbeat"
)

func Push(db *sql.DB, hb heartbeat.Heartbeat) error {
	statement := fmt.Sprintf("INSERT INTO heartbeats VALUES ('%d', '%s', '%d')", hb.UserId, hb.Location, hb.Timestamp)

	_, err := db.Exec(statement)

	if err != nil {
		log.Panic("Failed to push heartbeat to model", err.Error())
		return nil
	}

	return nil
}

func Get(db *sql.DB, id int) ([]heartbeat.HeartbeatTrack, error) {
	statement := fmt.Sprintf("SELECT * FROM tracks WHERE UserId = %d", id)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	hbtracks := []heartbeat.HeartbeatTrack{}

	var count = 0

	for rows.Next() {
		var hbt heartbeat.HeartbeatTrack
		if err := rows.Scan(&hbt.Uuid, &hbt.Heartbeats, &hbt.Starttime, &hbt.UserId); err != nil {
			return nil, err
		}

		hbtracks = append(hbtracks, hbt)
	}

	if count == 0 {
		return nil, sql.ErrNoRows
	}

	return hbtracks, nil
}

func DeleteHeartbeats(db *sql.DB, id int) error {
	return nil
}
