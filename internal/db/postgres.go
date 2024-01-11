package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func NewPostgresClient(DATASOURCE_NAME string) (*sql.DB, error) {
	db, connectionError := sql.Open("postgres", DATASOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
	fmt.Println("Connection to DB Successful")

	db.SetConnMaxLifetime(time.Minute * 3)

	return db, connectionError
}
