package gorp_migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/user"
	"tracker/proto/heartbeat"

	gorp "gopkg.in/gorp.v1"
)

func DropTables(db *sql.DB) {
	fmt.Println("Initalizing Tables in MSQL DB")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	dbmap.AddTableWithName(user.User{}, "users")
	dbmap.AddTableWithName(heartbeat.Heartbeat{}, "heartbeat")
	dbmap.AddTableWithName(heartbeat.HeartbeatTrack{}, "tracks")
	dbmap.DropTables()
	fmt.Println("Successfully Created Table")
}
