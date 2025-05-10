package models

type GetFlightsDTO struct {
	DepartDate  string `json:"depart_date,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Destination string `json:"destination,omitempty"`
}
