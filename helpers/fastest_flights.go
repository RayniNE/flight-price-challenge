package helpers

import (
	"sort"

	"github.com/raynine/flight-price-challenge/models"
)

func GetOrderedFlightByTime(flights []models.Flights) []models.Flights {
	orderedFlights := flights

	sort.Slice(orderedFlights, func(i, j int) bool {
		return orderedFlights[i].ArrivalTime < orderedFlights[j].ArrivalTime
	})

	return orderedFlights
}
