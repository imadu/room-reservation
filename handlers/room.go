package handlers

import (
	"github.com/avelino/slugify"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/oshosanya/room-reservation/constants"
	"github.com/oshosanya/room-reservation/db/repository"
	"github.com/oshosanya/room-reservation/request"
	"net/http"
)

func GetRooms(c echo.Context) error {
	return sendData(c, repository.GetRooms())
}

func CreateRoom(c echo.Context) error {
	r := new(request.Room)
	if err := c.Bind(r); err != nil {
		return sendError(c, "", "", "")
	}
	err := repository.CreateRoom(r.Name, slugify.Slugify(r.Name))
	if err != nil {
		log.Errorf("Could not create room: %+v", err)
		return sendError(c, "", "Could not create room", "")
	}
	return sendSuccess(c, nil)
}

func GetRoom(c echo.Context) error {
	slug := c.Param("slug")
	room, err := repository.GetRoom(slug)
	if err != nil {
		c.Response().WriteHeader(http.StatusNotFound)
		return sendError(c, "", constants.ROOM_DOES_NOT_EXIST, "")
	}
	return sendData(c, room)
}
