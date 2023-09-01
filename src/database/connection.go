package conn

import (
	"database/sql"
	"fmt"
	"log"
	// "os"
	"time"

	_ "github.com/lib/pq"
)

const (
	DRIVER_NAME     = "postgres"
	DATASOURCE_NAME = "postgres://postgres:G22mEyct@localhost:5432/portfoliodb?sslmode=disable"
)

var Db *sql.DB
var connectionError error

// var DATASOURCE_NAME = os.Getenv("DB_PASS")

func ConnectDB() {
	Db, connectionError = sql.Open(DRIVER_NAME, DATASOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
	fmt.Println("Connection to DB Successful")
	// getDbVersion()
	Db.SetConnMaxLifetime(time.Minute * 3)
}

func GetDb() *sql.DB {
	return Db
}

func getDbVersion() {
	// check db version
	rws, err := Db.Query("SELECT version();")
	if err != nil {
		log.Fatal("error getting version", err)
	}
	var ver string
	for rws.Next() {
		rws.Scan(&ver)
	}
	fmt.Printf("Current Database version is :: %s", ver)
}
