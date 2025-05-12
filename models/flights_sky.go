package models

import (
	"time"
)

type ResponseFlightsSky struct {
	Data    DataFlightsSky `json:"data"`
	Status  bool           `json:"status"`
	Message string         `json:"message"`
}

type DataFlightsSky struct {
	Context     ContextFlightsSky     `json:"context"`
	Itineraries []ItineraryFlightsSky `json:"itineraries"`
	Token       string                `json:"token"`
}

type ContextFlightsSky struct {
	Status       string `json:"status"`
	SessionID    string `json:"sessionId"`
	TotalResults int    `json:"totalResults"`
}

type ItineraryFlightsSky struct {
	ID    string          `json:"id"`
	Price PriceFlightsSky `json:"price"`
	Legs  []LegFlightsSky `json:"legs"`
}

type PriceFlightsSky struct {
	Raw             float64 `json:"raw"`
	Formatted       string  `json:"formatted"`
	PricingOptionID string  `json:"pricingOptionId"`
}

type LegFlightsSky struct {
	ID                string            `json:"id"`
	Origin            AirportFlightsSky `json:"origin"`
	Destination       AirportFlightsSky `json:"destination"`
	DurationInMinutes int               `json:"durationInMinutes"`
	StopCount         int               `json:"stopCount"`
	IsSmallestStops   bool              `json:"isSmallestStops"`
	Departure         string            `json:"departure"`
	Arrival           string            `json:"arrival"`
	TimeDeltaInDays   int               `json:"timeDeltaInDays"`
}

type AirportFlightsSky struct {
	ID            string `json:"id"`
	EntityID      string `json:"entityId"`
	Name          string `json:"name"`
	DisplayCode   string `json:"displayCode"`
	City          string `json:"city"`
	Country       string `json:"country"`
	IsHighlighted bool   `json:"isHighlighted"`
}

func (model *ResponseFlightsSky) MapPriceLineToModel() []Flights {
	flights := []Flights{}

	if model.Data.Itineraries == nil {
		return flights
	}

	for _, itinerary := range model.Data.Itineraries {
		if len(itinerary.Legs) > 0 {
			leg := itinerary.Legs[0]

			arrival, _ := time.Parse(DATE_FORMAT, leg.Arrival)
			departure, _ := time.Parse(DATE_FORMAT, leg.Departure)

			flight := Flights{
				Price:           itinerary.Price.Raw,
				OriginName:      leg.Origin.Name,
				ArrivalTime:     arrival,
				DepartureTime:   departure,
				DestinationName: leg.Destination.Name,
			}

			flights = append(flights, flight)
		}
	}

	return flights
}
