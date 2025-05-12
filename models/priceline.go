package models

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
