package models

import (
	"time"
)

type Reservation struct {
	Model
	ReservistID string    `json:"reservist_id"`
	RoomID      uint      `json:"room_id"`
	Room        Room      `gorm:"PRELOAD:true" json:"room"`
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
}
