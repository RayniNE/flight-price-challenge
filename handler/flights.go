package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
