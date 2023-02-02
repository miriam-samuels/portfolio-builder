package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

const (
	DRIVER_NAME     = "mysql"
	DATASOURCE_NAME = "root:password@/portfolio"
)

var PortfolioDb *sql.DB
var connectionError error

func ConnectDB() {
	PortfolioDb, connectionError = sql.Open(DRIVER_NAME, DATASOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
}

func GetDb(w http.ResponseWriter, r *http.Request) {
	//perform select query to get db
	rows, err := PortfolioDb.Query("SELECT DATABASE() AS db")
	if err != nil {
		log.Fatal("error executing db select query :: ", err)
	}

	var rec string
	for rows.Next() {
		rows.Scan(&rec)
	}

	fmt.Fprintf(w, "Current Database is :: %s", rec)
}
