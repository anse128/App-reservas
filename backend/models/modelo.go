// models/modelo.go

package models

import "time"

type Modelo struct {
	ID            uint      `gorm:"primaryKey"`
	Nombre        string    `gorm:"not null"`
	Descripcion   string    `gorm:"not null"`
	FechaCreacion time.Time `gorm:"autoCreateTime"`
}
