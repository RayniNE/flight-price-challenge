package models

import "time"

// Agoda contiene la sección de datos
type AgodaResponse struct {
	Data AgodaData `json:"data"`
}

// AgodaData contiene la lista de bundles (ofertas/paquetes)
type AgodaData struct {
	Bundles []Bundle `json:"bundles"`
}

// Bundle representa una oferta individual
type Bundle struct {
	BundlePrice   []BundlePriceEntry `json:"bundlePrice"`
	OutboundSlice OutboundSlice      `json:"outboundSlice"`
}

// BundlePriceEntry contiene la información del precio del bundle
type BundlePriceEntry struct {
	Price PriceInfo `json:"price"`
}

// PriceInfo contiene detalles del precio, incluyendo la moneda USD
type PriceInfo struct {
	USD USD `json:"usd"`
}

// USD contiene la visualización del precio en dólares
type USD struct {
	Display PriceDisplay `json:"display"`
}

// PriceDisplay muestra el precio por reserva
type PriceDisplay struct {
	PerBook PerBook `json:"perBook"`
}

// PerBook contiene el precio total con todo incluido
type PerBook struct {
	AllInclusive float64 `json:"allInclusive"`
}

// Itinerary representa un itinerario de vuelo
type Itinerary struct {
	OutboundSlice OutboundSlice `json:"outboundSlice"`
}

// OutboundSlice representa el tramo de ida del vuelo
type OutboundSlice struct {
	Segments []Segment `json:"segments"`
}

// Segment representa un segmento individual del vuelo (puede haber múltiples para conexiones)
type Segment struct {
	ArrivalDateTime string         `json:"arrivalDateTime"`
	DepartDateTime  string         `json:"departDateTime"`
	AirportContent  AirportContent `json:"airportContent"`
}

// AirportContent contiene información detallada sobre los aeropuertos de salida y llegada
type AirportContent struct {
	DepartureCityID      int    `json:"departureCityId"`
	DepartureCityName    string `json:"departureCityName"`
	DepartureCountryID   int    `json:"departureCountryId"`
	DepartureCountryName string `json:"departureCountryName"`
	DepartureAirportName string `json:"departureAirportName"`
	ArrivalCityID        int    `json:"arrivalCityId"`
	ArrivalCityName      string `json:"arrivalCityName"`
	ArrivalCountryID     int    `json:"arrivalCountryId"`
	ArrivalCountryName   string `json:"arrivalCountryName"`
	ArrivalAirportName   string `json:"arrivalAirportName"`
}

func (model *AgodaResponse) MapPriceLineToModel() []Flights {
	flights := []Flights{}

	if model.Data.Bundles == nil {
		return flights
	}

	for _, bundle := range model.Data.Bundles {
		flight := Flights{}

		if len(bundle.BundlePrice) > 0 && bundle.BundlePrice[0].Price.USD.Display.PerBook.AllInclusive > 0 {
			flight.Price = bundle.BundlePrice[0].Price.USD.Display.PerBook.AllInclusive
		}

		if len(bundle.OutboundSlice.Segments) > 0 {
			segment := bundle.OutboundSlice.Segments[0]

			arrival, _ := time.Parse(DATE_FORMAT, segment.ArrivalDateTime)
			departure, _ := time.Parse(DATE_FORMAT, segment.DepartDateTime)
			flight.ArrivalTime = arrival
			flight.DepartureTime = departure
			flight.OriginName = segment.AirportContent.DepartureAirportName
			flight.DestinationName = segment.AirportContent.ArrivalAirportName
		}
		flights = append(flights, flight)
	}

	return flights
}
