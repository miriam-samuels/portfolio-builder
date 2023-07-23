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
	// CurrentDb()
	fmt.Println("Connection to DB Successful")
	Db.SetConnMaxLifetime(time.Minute * 3)

	createTable()
}

func createTable() {
	var exists bool
	err := Db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)", "users").Scan(&exists)
	if err != nil {
		_, err := Db.Exec(`
		CREATE TABLE users(
			id VARCHAR(40) PRIMARY KEY,
			username VARCHAR(30) NOT NULL ,
			password VARCHAR(100) NOT NULL,
			email VARCHAR(45) DEFAULT '',
			first_name VARCHAR(45) DEFAULT '',
			last_name VARCHAR(45) DEFAULT '',
			phone VARCHAR(20) DEFAULT '',
			github VARCHAR(45) DEFAULT '',
			medium VARCHAR(45) DEFAULT '',
			twitter VARCHAR(45) DEFAULT '',
			linkedin VARCHAR(45) DEFAULT '',
			objective VARCHAR(400) DEFAULT '',
			tagline VARCHAR(150) DEFAULT '',
			skills JSON,
			projects JSON,
			theme VARCHAR(30) DEFAULT ''
		);
		`)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Table created successfully.")
	}

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
