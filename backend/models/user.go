// models/usuario.go

package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre   string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Reservas []Reserva
}
