package main

import (
	"database/sql"
	"log"
	"net/http"
	"proyectofinalweb/backend/handlers" // Ajusta la ruta según la estructura de tu proyecto

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=secret dbname=car_rental sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/api/register", handlers.RegisterHandler(db)).Methods("POST")
	router.HandleFunc("/api/login", handlers.LoginHandler(db)).Methods("POST")
	router.HandleFunc("/api/cars", handlers.SearchCarsHandler(db)).Methods("GET")
	router.HandleFunc("/api/reserve", handlers.ReserveCarHandler(db)).Methods("POST")
	router.HandleFunc("/api/report", handlers.ReportHandler(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
