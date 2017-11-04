package gorp_migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"tracker/proto/heartbeat"

	gorp "gopkg.in/gorp.v1"
)

type MigratorInterface interface {
	CreateTables()
}

type Migrator struct {
}

func CreateTables(db *sql.DB) {
	fmt.Println("Initalizing Tables in MSQL DB")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	//	dbmap.AddTableWithName(user.User{}, "users").SetKeys(true, "id")
	dbmap.AddTableWithName(heartbeat.Heartbeat{}, "heartbeat").SetKeys(true, "Timestamp")
	dbmap.CreateTablesIfNotExists()
	fmt.Println("Successfully Created Table")
}
