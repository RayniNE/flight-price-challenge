package helpers

import (
	"sort"

	"github.com/raynine/flight-price-challenge/models"
)

func GetOrderedFlightByTime(flights []models.Flights) []models.Flights {
	orderedFlights := flights

	sort.Slice(orderedFlights, func(i, j int) bool {
		firstTime := orderedFlights[i].DepartureTime.Sub(orderedFlights[i].ArrivalTime)
		secondTime := orderedFlights[j].DepartureTime.Sub(orderedFlights[j].ArrivalTime)
		return firstTime < secondTime
	})

	return orderedFlights
}

func GetOrderedFlightByPrice(flights []models.Flights) []models.Flights {
	orderedFlights := flights

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})

	return orderedFlights
}
