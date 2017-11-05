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
	dbmap.AddTableWithName(heartbeat.Heartbeat{}, "heartbeat")
	dbmap.AddTableWithName(heartbeat.HeartbeatTrack{}, "tracks").SetKeys(false, "Uuid")
	dbmap.CreateTablesIfNotExists()
	if err := updateToType(db, "tracks", "Heartbeats", "BLOB"); err != nil {
		log.Fatal("Failed to create track table. Exiting")
	}

	if err := addPrimaryIDColumn(db, "heartbeat"); err != nil {
		log.Fatal("Failed to add primary key column")
	}

	fmt.Println("Successfully Created Table")
}

//These are necessary because I don't want a complete ORM. There are a few fields that I don't want to maintain.
func addPrimaryIDColumn(db *sql.DB, table string) error {
	statement := fmt.Sprintf("ALTER TABLE %s ADD ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY", table)

	_, err := db.Exec(statement)
	if err != nil {
		fmt.Println("Error updating primary key", err.Error())
	}
	return err
}

//NOTE: Run on empty table to avoid failure
func updateToType(db *sql.DB, table, column, newtype string) error {
	statement := fmt.Sprintf("ALTER TABLE %s MODIFY %s %s", table, column, newtype)

	_, err := db.Exec(statement)
	if err != nil {
		fmt.Println("Error handling update to column")
	}

	return err
}
