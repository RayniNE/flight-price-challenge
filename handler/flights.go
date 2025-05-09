package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raynine/flight-price-challenge/helpers"
)

type FlightsHandler struct{}

func NewFlightsHandler() *FlightsHandler {
	return &FlightsHandler{}
}

func (handler *FlightsHandler) GetFlights(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"hello": "world",
	})
}

func (handler *FlightsHandler) GetAutoComplete(c echo.Context) error {
	response, err := helpers.GetSkyScannerFlights()
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, response)
	return nil
}
