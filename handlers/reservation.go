package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/oshosanya/room-reservation/constants"
	"github.com/oshosanya/room-reservation/db/repository"
	"github.com/oshosanya/room-reservation/request"
	"time"
)

func CreateReservation(c echo.Context) error {
	roomSlug := c.Param("slug")
	r := new(request.Reservation)
	if err := c.Bind(r); err != nil {
		log.Errorf("Could not bind request to struct: %+v", err)
		return sendError(c, "", "", "")
	}
	room, err := repository.GetRoom(roomSlug)
	if err != nil {
		return sendError(c, "", constants.ROOM_DOES_NOT_EXIST, "")
	}
	reservations, err := repository.GetOpenReservationsByRoom(room.ID, r.From)
	if len(reservations) > 0 {
		return sendError(c, "", fmt.Sprintf("Cannot reserve room. Last reservation is for %s", reservations[0].To), "")
	}

	err = repository.CreateReservation(r.ReservistID, room.Model.ID, r.From, r.To)
	if err != nil {
		return sendError(c, "", "", "")
	}

	return sendSuccess(c, nil)
}

func GetReservationByRoom(c echo.Context) error {
	roomSlug := c.Param("slug")
	room, err := repository.GetRoom(roomSlug)
	if err != nil {
		return sendError(c, "", constants.ROOM_DOES_NOT_EXIST, "")
	}
	reservations, err := repository.GetOpenReservationsByRoom(room.Model.ID, time.Now())
	if err != nil {
		log.Errorf("Could not get reservations for room: %+v", err)
		return sendError(c, "", constants.AN_ERROR_OCCURED, "")
	}
	return sendData(c, reservations)
}

func GetReservations(c echo.Context) error {
	reservations, err := repository.GetOpenReservations()
	if err != nil {
		log.Errorf("Could not get reservations for room: %+v", err)
		return sendError(c, "", constants.AN_ERROR_OCCURED, "")
	}
	return sendData(c, reservations)
}
