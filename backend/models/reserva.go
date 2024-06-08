// models/reserva.go

package models

import (
	"time"

	"gorm.io/gorm"
)

type Reserva struct {
	gorm.Model
	UsuarioID uint
	CarroID   uint
	Fecha     time.Time `gorm:"autoCreateTime"`
}
