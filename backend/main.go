package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	`C:/Users/bra/OneDrive - Universidad de Antioquia/Escritorio/proyectofinalweb/backend/handlers`
	`proyectofinalweb/backend/handlers` // Importa el paquete que contiene las funciones de manejo
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=secret dbname=car_rental sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/api/register", handlers.RegisterHandler(db)).Methods("POST") // Usa las funciones calificadas por el paquete
	router.HandleFunc("/api/login", handlers.LoginHandler(db)).Methods("POST")       // handlers.RegisterHandler(db) en lugar de RegisterHandler(db)
	router.HandleFunc("/api/cars", handlers.SearchCarsHandler(db)).Methods("GET")
	router.HandleFunc("/api/reserve", handlers.ReserveCarHandler(db)).Methods("POST")
	router.HandleFunc("/api/report", handlers.ReportHandler(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
