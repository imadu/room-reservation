package repository

import (
	"github.com/oshosanya/room-reservation/db/models"
	"time"
)

func CreateReservation(reservistID string, roomID uint, from time.Time, to time.Time) error {
	newReservation := models.Reservation{}
	newReservation.ReservistID = reservistID
	newReservation.RoomID = roomID
	newReservation.From = from
	newReservation.To = to
	return database.Create(&newReservation).Error
}

func GetRervations() ([]models.Reservation, error) {
	var reservation []models.Reservation
	if err := database.Find(&reservation).Error; err != nil {
		return reservation, err
	}
	return reservation, nil
}

func GetOpenReservations() ([]models.Reservation, error) {
	var reservation []models.Reservation
	if err := database.Where("`to` > ?", time.Now()).Order("`to` desc").Find(&reservation).Error; err != nil {
		return reservation, err
	}
	return reservation, nil
}

func GetReservationsByRoom(roomID int) ([]models.Reservation, error) {
	var reservation []models.Reservation
	if err := database.Where("room_id = ?", roomID).Find(&reservation).Error; err != nil {
		return reservation, err
	}
	return reservation, nil
}

func GetOpenReservationsByRoom(roomID uint, from time.Time) ([]models.Reservation, error) {
	var reservation []models.Reservation
	if err := database.Where("room_id = ? AND `to` > ?", roomID, from).Order("`to` desc").Find(&reservation).Error; err != nil {
		return reservation, err
	}
	return reservation, nil
}
