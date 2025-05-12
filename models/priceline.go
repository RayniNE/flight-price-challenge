package models

import "time"

const DATE_FORMAT = "2006-01-02T15:04:05"

type PricelineResponse struct {
	Data PricelineData `json:"data"`
}

type PricelineData struct {
	Listings []Listing `json:"listings"`
}

type Listing struct {
	TotalPriceWithDecimal TotalPrice `json:"totalPriceWithDecimal"`
	Slices                []Slice    `json:"slices"`
}

type TotalPrice struct {
	Price float64 `json:"price"`
}

type Slice struct {
	Segments []PriceLineSegment `json:"segments"`
}

type PriceLineSegment struct {
	DepartInfo  FlightInfoPoint `json:"departInfo"`
	ArrivalInfo FlightInfoPoint `json:"arrivalInfo"`
}

type FlightInfoPoint struct {
	Airport AirportDetails `json:"airport"`
	Time    TimeDetails    `json:"time"`
}

type AirportDetails struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type TimeDetails struct {
	DateTime string `json:"dateTime"`
}

// --- Estructura para los datos extraídos de Priceline ---

// PricelineAirportContent agrupa la información de los aeropuertos para Priceline
type PricelineAirportContent struct {
	DepartureAirportName string
	DepartureAirportCode string
	ArrivalAirportName   string
	ArrivalAirportCode   string
}

// ExtractedPricelineInfo contiene los campos solicitados para cada opción de vuelo de Priceline
type ExtractedPricelineInfo struct {
	Price           float64
	ArrivalDateTime string
	DepartDateTime  string
	AirportInfo     PricelineAirportContent // Adaptado para la estructura de Priceline
}

func (model *PricelineResponse) MapPriceLineToModel() []Flights {
	// Recorrer itinerarios y extraer información
	flights := []Flights{}

	if model.Data.Listings == nil {
		return flights // No hay listings para procesar
	}

	for _, listing := range model.Data.Listings {
		flight := Flights{}

		flight.Price = listing.TotalPriceWithDecimal.Price

		// Se asume que el primer slice y el primer segmento son los relevantes.
		if len(listing.Slices) > 0 {
			slice := listing.Slices[0]
			if len(slice.Segments) > 0 {
				segment := slice.Segments[0]
				arrival, _ := time.Parse(DATE_FORMAT, segment.ArrivalInfo.Time.DateTime)
				departure, _ := time.Parse(DATE_FORMAT, segment.DepartInfo.Time.DateTime)
				flight.DepartureTime = departure
				flight.ArrivalTime = arrival
				flight.DestinationName = segment.ArrivalInfo.Airport.Name
				flight.OriginName = segment.DepartInfo.Airport.Name
			}
		}
		flights = append(flights, flight)
	}

	return flights
}
