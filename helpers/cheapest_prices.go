package helpers

import (
	"sort"
	"time"

	"github.com/raynine/flight-price-challenge/models"
)

const DATE_FORMAT = "2006-01-02T15:04:05"

func GetOrderedFlightsSkyByPrice(model *models.ResponseFlightsSky) []models.Flights {
	flights := []models.Flights{}

	if model.Data.Itineraries == nil {
		return flights
	}

	for _, itinerary := range model.Data.Itineraries {
		if len(itinerary.Legs) > 0 {
			leg := itinerary.Legs[0]

			arrival, _ := time.Parse(DATE_FORMAT, leg.Arrival)
			departure, _ := time.Parse(DATE_FORMAT, leg.Departure)

			flight := models.Flights{
				Price:           itinerary.Price.Raw,
				OriginName:      leg.Origin.Name,
				ArrivalTime:     arrival,
				DepartureTime:   departure,
				DestinationName: leg.Destination.Name,
			}

			flights = append(flights, flight)
		}
	}

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})

	return flights
}

func GetOrderedAgodaFlightsByPrice(model *models.AgodaResponse) []models.Flights {
	flights := []models.Flights{}

	if model.Data.Bundles == nil {
		return flights
	}

	for _, bundle := range model.Data.Bundles {
		flight := models.Flights{}

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

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})

	return flights
}

func GetOrderedPriceLineFlightsByPrice(model *models.PricelineResponse) []models.Flights {
	// Recorrer itinerarios y extraer información
	flights := []models.Flights{}

	if model.Data.Listings == nil {
		return flights // No hay listings para procesar
	}

	for _, listing := range model.Data.Listings {
		flight := models.Flights{}

		flight.Price = listing.TotalPriceWithDecimal.Price

		// Se asume que el primer slice y el primer segmento son los relevantes.
		if len(listing.Slices) > 0 {
			slice := listing.Slices[0]
			if len(slice.Segments) > 0 {
				segment := slice.Segments[0]
				arrival, _ := time.Parse(DATE_FORMAT, segment.ArrivalInfo.Time.DateTime)
				departure, _ := time.Parse(DATE_FORMAT, segment.DepartInfo.Time.DateTime)
				flight.DepartureTime = arrival
				flight.ArrivalTime = departure
				flight.DestinationName = segment.ArrivalInfo.Airport.Name
				flight.OriginName = segment.DepartInfo.Airport.Name
			}
		}
		flights = append(flights, flight)
	}

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})

	return flights
}
