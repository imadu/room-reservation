package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/oshosanya/room-reservation/db"
)

var database *gorm.DB

func ConnectToDB() {
	database = db.GetDB()
	database.LogMode(true)
	database = database.Set("gorm:auto_preload", true)
}
