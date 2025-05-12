package models

type GetFlightsDTO struct {
	DepartDate  string `json:"depart_date,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Destination string `json:"destination,omitempty"`
}

type Flights struct {
	Price           float64 `json:"price"`
	OriginCode      string  `json:"origin_code"`
	OriginName      string  `json:"origin_name"`
	DestinationCode string  `json:"destination_code"`
	DestinationName string  `json:"destination_name"`
	DepartureTime   string  `json:"departure_time"`
	ArrivalTime     string  `json:"arrival_time"`
}
