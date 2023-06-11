package conn

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	DRIVER_NAME = "postgres"
)

var Db *sql.DB
var connectionError error
var DATASOURCE_NAME = os.Getenv("DB_PASS")

func ConnectDB() {
	Db, connectionError = sql.Open(DRIVER_NAME, DATASOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
	CurrentDb()
	fmt.Println("Connection to DB Successful")
	Db.SetConnMaxLifetime(time.Minute * 3)
}

func GetDb() *sql.DB {
	return Db
}

func CurrentDb() {
	// perform select query to get db
	rows, err := Db.Query("SELECT DATABASE() AS db")
	if err != nil {
		log.Fatal("error executing db select query :: ", err)
	}

	var rec string
	for rows.Next() {
		rows.Scan(&rec)
	}

	fmt.Printf("Current Database is :: %s", rec)
}
