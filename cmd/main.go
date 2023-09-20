package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	conn "github.com/miriam-samuels/src/database"
	"github.com/miriam-samuels/src/routes/v1"
)

const (
	CONN_HOST = "portfolio-builder-qndq.onrender.com"
	CONN_PORT = "5432"
)

func init() {
	conn.ConnectDB()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = CONN_PORT
	}

	router := mux.NewRouter().StrictSlash(true)
	v1 := router.PathPrefix("/").Subrouter()
	// v2 := router.PathPrefix("/v2").Subrouter()

	// http.Handle("/", router)

	routes.RoutesV1(v1)

	defer conn.Db.Close()

	err := http.ListenAndServe(CONN_HOST, router)
	if err != nil {
		log.Fatal(err)
	}
}
