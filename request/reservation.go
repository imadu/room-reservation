package request

import "time"

type Reservation struct {
	ReservistID string    `json:"reservist_id"`
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
}
