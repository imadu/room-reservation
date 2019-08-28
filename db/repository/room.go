package repository

import (
	"github.com/oshosanya/room-reservation/db/models"
	"github.com/oshosanya/room-reservation/util"
	"github.com/pkg/errors"
)

func GetRooms() []models.Room {
	var rooms []models.Room
	database.Find(&rooms)
	return rooms
}

func CreateRoom(name string, slug string) error {
	existingRoom := models.Room{}
	newRoom := models.Room{}
	// Check if room with slug does not exist
	if err := database.Where("slug = ?", slug).First(&existingRoom).Error; err != nil {
		newRoom.Name = name
		newRoom.Slug = slug
		return database.Create(&newRoom).Error
	}

	// Append suffix to slug if slug has been previously used
	newRoom.Name = name
	newRoom.Slug = slug + "-" + util.RandStringBytes(2)

	return database.Create(&newRoom).Error
}

func GetRoom(slug string) (models.Room, error) {
	room := models.Room{}
	if err := database.Where("slug = ?", slug).First(&room).Error; err != nil {
		return room, errors.New("Room does not exist")
	}
	return room, nil
}
