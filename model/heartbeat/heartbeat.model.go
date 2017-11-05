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

func Get(db *sql.DB, id int) (heartbeat.Heartbeat, error) {
	statement := fmt.Sprintf("SELECT * FROM heartbeats WHERE uid = %d", id)
	fmt.Println("Statment", statement)
	return heartbeat.Heartbeat{}, nil
}

func DeleteHeartbeats(db *sql.DB, id int) error {
	return nil
}
