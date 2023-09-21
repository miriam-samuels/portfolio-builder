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
	DRIVER_NAME = "postgres"
	DATASOURCE_NAME = "postgres://root:7hPXtNDOLj5LMhGbdJlvIrMjmE7V14na@dpg-ck5nrl8s0i2c73bai7s0-a.oregon-postgres.render.com/portfolio_2p6g?sslmode=disable"

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
