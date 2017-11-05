package gorp_migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"tracker/proto/heartbeat"
	"tracker/proto/user"

	gorp "gopkg.in/gorp.v1"
)

func CreateTables(db *sql.DB) {
	fmt.Println("Initalizing Tables in MSQL DB")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "Tracker:", log.Lmicroseconds))
	dbmap.AddTableWithName(user.User{}, "users").SetKeys(true, "UserId")
	dbmap.AddTableWithName(heartbeat.Heartbeat{}, "heartbeat").SetKeys(true, "Timestamp")
	dbmap.AddTableWithName(heartbeat.HeartbeatTrack{}, "tracks").SetKeys(false, "Uuid")
	dbmap.CreateTablesIfNotExists()
	if err := updateToType(db, "tracks", "Heartbeats", "BLOB"); err != nil {
		log.Fatal("Failed to create track table. Exiting")
	}
	fmt.Println("Successfully Created Table")
}

//NOTE: Can only work on an empty table
func updateToType(db *sql.DB, table, column, newtype string) error {
	statement := fmt.Sprintf("ALTER TABLE %s MODIFY %s %s", table, column, newtype)
	fmt.Println(statement)
	_, err := db.Exec(statement)
	if err != nil {
		fmt.Println("Error handling update to column")
	}

	return err
}
