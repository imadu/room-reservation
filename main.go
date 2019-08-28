package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/oshosanya/room-reservation/db/repository"
	"github.com/oshosanya/room-reservation/handlers"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/subosito/gotenv"
)

func main() {
	err := gotenv.Load()
	if err != nil {
		log.Fatalf("Could not load env file: %+v", err)
	}
	appPort := os.Getenv("APP_PORT")
	repository.ConnectToDB()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/room", handlers.GetRooms)
	e.POST("/room", handlers.CreateRoom)
	e.GET("/room/:slug", handlers.GetRoom)
	e.POST("/room/:slug/reserve", handlers.CreateReservation)
	e.GET("/room/:slug/reservations", handlers.GetReservationByRoom)
	e.GET("/reservations", handlers.GetReservations)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", appPort)))
}
