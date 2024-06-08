package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Car struct {
	ID          int    `json:"id"`
	Marca       string `json:"marca"`
	Modelo      string `json:"modelo"`
	Combustible string `json:"combustible"`
}

type Reservation struct {
	ID     int `json:"id"`
	CarID  int `json:"car_id"`
	UserID int `json:"user_id"`
}

type Report struct {
	Reservations []Reservation `json:"reservations"`
	TotalCost    float64       `json:"total_cost"`
}
