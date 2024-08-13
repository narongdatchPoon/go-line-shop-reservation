package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	ID              uint   `json:"id" gorm:"primary_key"`
	Title           string `json:"title"`
	Time            string `json:"time"`
	ReservationDate time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
