package models

import "time"

type GetFlightsDTO struct {
	DepartDate    time.Time `json:"depart_date,omitempty"`
	ReturnDate    time.Time `json:"return_date,omitempty"`
	Origin        string    `json:"origin,omitempty"`
	OriginID      string    `json:"origin_id,omitempty"`
	Destination   string    `json:"destination,omitempty"`
	DestinationID string    `json:"destination_id,omitempty"`
}

type GetAutoCompleteResponse struct {
	InputSuggest []InputSuggest `json:"inputSuggest,omitempty"`
}

type InputSuggest struct {
	Navigation   Navigation   `json:"navigation,omitempty"`
	Presentation Presentation `json:"presentation,omitempty"`
}

type Presentation struct {
	Subtitle        string `json:"subtitle,omitempty"`
	SuggestionTitle string `json:"suggestionTitle,omitempty"`
	Title           string `json:"title,omitempty"`
}

type Navigation struct {
	RelevantFlightParams RelevantFlightParams `json:"relevantFlightParams,omitempty"`
}

type RelevantFlightParams struct {
	EntityID        string `json:"entityId,omitempty"`
	LocalizedName   string `json:"localizedName,omitempty"`
	FlightPlaceType string `json:"flightPlaceType,omitempty"`
	SkyID           string `json:"skyId,omitempty"`
}
