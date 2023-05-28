package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	conn "github.com/miriam-samuels/src/database"
	"github.com/miriam-samuels/src/routes/v1"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3030"
)

func init() {
	conn.ConnectDB()
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	v1 := router.PathPrefix("/").Subrouter()
	// v2 := router.PathPrefix("/v2").Subrouter()

	// http.Handle("/", router)

	routes.RoutesV1(v1)

	defer conn.Db.Close()

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal(err)
	}
}
