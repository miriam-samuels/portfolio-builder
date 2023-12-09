package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/miriam-samuels/portfolio-builder/internal/db"
	"github.com/miriam-samuels/portfolio-builder/internal/routes/v1"
	"github.com/rs/cors"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3000"
)

func init() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Connect Database
	client, err := db.NewPostgresClient(os.Getenv("PORTFOLIO_DB_DATASOURCE_URI"))
	if err != nil {
		log.Fatal("error connecting to database :: ", err)
	}

	// set db to created client
	db.Portfolio = client
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

	defer db.Portfolio.Close()

	//  cross origin
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		// Debug:            true,
	}).Handler(router)

	// add more configurations to server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	// start server
	fmt.Println("starting server on port :: " + port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
