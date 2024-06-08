package handlers

import (
	"encoding/json"
	"net/http"
	"proyectofinalweb/backend/models"
	"database/sql"
)

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		// Lógica para registrar usuario
		// ...

		w.WriteHeader(http.StatusCreated)
	}
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		// Lógica para autenticar usuario
		// ...

		w.WriteHeader(http.StatusOK)
	}
}

func SearchCarsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Lógica para buscar autos
		// ...

		cars := []models.Car{
			{ID: 1, Marca: "Toyota", Modelo: "Corolla", Combustible: "Gasolina"},
			{ID: 2, Marca: "Honda", Modelo: "Civic", Combustible: "Gasolina"},
		}

		json.NewEncoder(w).Encode(cars)
	}
}
