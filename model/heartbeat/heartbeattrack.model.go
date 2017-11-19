package heartbeat

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"tracker/proto/heartbeat"
	puuid "tracker/proto/uuid"
)

//Theres a dependency between a heartbeat and a heartbeat track. You cannot have a heartbeat point to a heartbeat tracka and the heartbeat track not exist.

func Push(db *sql.DB, hb heartbeat.Heartbeat) error {

	statement := fmt.Sprintf("INSERT INTO heartbeat (UserId, Location, Longitude, Latitude, Timestamp, Altitude) VALUES ('%d', '%s', '%v', '%v', '%v', '%v')", hb.UserId, hb.Location, hb.Longitude, hb.Latitude, hb.Timestamp, hb.Altitude)

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
		var endtime sql.NullInt64

		if err := rows.Scan(&uid, &hbs, &hbt.Starttime, &endtime, &hbt.UserId); err != nil {
			return nil, err
		}

		var t puuid.UUID
		t.Value = uid
		hbt.Uuid = &t
		addIfNotNil(hbt.Endtime, endtime)

		var hbs2 = make([]*heartbeat.Heartbeat, 1, 1)
		//get points
		hbs2, err := GetHeartbeatsForTrack(db, hbt)
		if err != nil {
			log.Panic("Failed to fetch heartbeats: ", err)
		}
		hbt.Heartbeats = hbs2
		hbtracks = append(hbtracks, hbt)
	}

	if len(hbtracks) == 0 {
		return nil, sql.ErrNoRows
	}
	fmt.Println("Returnning ", len(hbtracks), " heartbeat tracks")
	fmt.Println("Running ", statement)
	return hbtracks, nil
}

func addIfNotNil(place, vv interface{}) {
	scan := reflect.TypeOf((*sql.Scanner)(nil)).Elem()
	uses := !reflect.PtrTo(reflect.TypeOf(vv)).Implements(scan)
	if uses {
		place = vv
	}
}
func DeleteHeartbeats(db *sql.DB, id int) error {
	return nil
}

//starttime and endtime.
func GetHeartbeatsForTrack(db *sql.DB, hbt heartbeat.HeartbeatTrack) ([]*heartbeat.Heartbeat, error) {

	endtime := hbt.Endtime
	var statement string
	var location sql.NullString
	var altitude sql.NullFloat64
	fmt.Println("endtime is ", endtime)
	//needs to limit by endtime as well.
	if endtime == 0 {
		statement = fmt.Sprintf("SELECT * FROM heartbeat WHERE UserId = %v AND Timestamp > %v ", hbt.UserId, hbt.Starttime)
	} else {
		statement = fmt.Sprintf("SELECT * FROM heartbeat WHERE UserId = %v AND Timestamp >= %v AND Timestamp <= %v ", hbt.UserId, hbt.Starttime, hbt.Endtime)
	}

	fmt.Println("Running statment ", statement)
	rows, err := db.Query(statement)
	var id int64
	var heartbeats []*heartbeat.Heartbeat
	if err != nil {
		log.Panic("Failed to retrieve heartbeats for the user", err, "\nRunning statement", statement)
	}

	for rows.Next() {
		var hb heartbeat.Heartbeat
		if err := rows.Scan(&hb.Timestamp, &location, &hb.UserId, &hb.Longitude, &hb.Latitude, &altitude, &id); err != nil {
			return nil, err
		}

		addIfNotNil(hb.Altitude, altitude)

		heartbeats = append(heartbeats, &hb)
	}
	return heartbeats, nil
}
