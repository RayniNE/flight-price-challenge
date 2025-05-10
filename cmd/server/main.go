package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/raynine/flight-price-challenge/handler"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := echo.New()

	flightsHandler := handler.NewFlightsHandler()

	e.GET("/flights/search", flightsHandler.GetFlights)
	e.GET("/airports", flightsHandler.GetAirports)

	e.Logger.Fatal(e.Start(":3000"))
}
