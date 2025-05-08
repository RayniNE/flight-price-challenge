package main

import (
	"github.com/labstack/echo/v4"
	"github.com/raynine/flight-price-challenge/handler"
)

func main() {
	e := echo.New()

	flightsHandler := handler.NewFlightsHandler()

	e.GET("/", flightsHandler.GetFlights)

	e.Logger.Fatal(e.Start(":3000"))
}
