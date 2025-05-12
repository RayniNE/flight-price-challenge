package models

import "time"

type GetFlightsDTO struct {
	DepartDate  string `json:"depart_date,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Destination string `json:"destination,omitempty"`
}

type Flights struct {
	Price           float64   `json:"price"`
	OriginCode      string    `json:"origin_code"`
	OriginName      string    `json:"origin_name"`
	DestinationCode string    `json:"destination_code"`
	DestinationName string    `json:"destination_name"`
	DepartureTime   time.Time `json:"departure_time,omitempty"`
	ArrivalTime     time.Time `json:"arrival_time,omitempty"`
}

type FlightsResponse struct {
	CheapestFlights []Flights `json:"cheapest_flights"`
	FastestFlights  []Flights `json:"fastest_flights"`
	OtherFlights    []Flights `json:"other_flights"`
}
