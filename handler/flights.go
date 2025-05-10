package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/raynine/flight-price-challenge/helpers"
	"github.com/raynine/flight-price-challenge/models"
)

type FlightsHandler struct{}

func NewFlightsHandler() *FlightsHandler {
	return &FlightsHandler{}
}

const NO_LOCATION = ""

func (handler *FlightsHandler) GetFlights(c echo.Context) error {
	origin := c.QueryParam("origin")
	destination := c.QueryParam("destination")
	date := c.QueryParam("date")
	dto := models.GetFlightsDTO{
		DepartDate:  date,
		Origin:      origin,
		Destination: destination,
	}

	response, err := helpers.GetPricelineFlights(dto)
	if err != nil {
		fmt.Printf("An error ocurred while opening json file: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": fmt.Sprintf("An error ocurred while opening json file: %s", err.Error()),
		})
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func (handler *FlightsHandler) GetAirports(c echo.Context) error {
	file, err := os.Open("airports.json")
	if err != nil {
		fmt.Printf("An error ocurred while opening json file: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": fmt.Sprintf("An error ocurred while opening json file: %s", err.Error()),
		})
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("An error ocurred while converting to bytes: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": fmt.Sprintf("An error ocurred while converting to bytes: %s", err.Error()),
		})
	}

	result := []map[string]any{}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		fmt.Printf("An error ocurred while converting to map: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": fmt.Sprintf("An error ocurred while converting to map: %s", err.Error()),
		})
	}

	c.JSON(http.StatusOK, result)
	return nil
}
