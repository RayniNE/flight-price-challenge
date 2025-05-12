package helpers

import (
	"sort"

	"github.com/raynine/flight-price-challenge/models"
)

func GetOrderedFlightsSkyByPrice(model *models.ResponseFlightsSky) []models.Flights {
	flights := []models.Flights{}

	for _, itinerary := range model.Data.Itineraries {
		if len(itinerary.Legs) > 0 {
			leg := itinerary.Legs[0]

			flight := models.Flights{
				Price:           itinerary.Price.Raw,
				OriginCode:      leg.Origin.Name,
				DestinationCode: leg.Destination.Name,
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
			flight.ArrivalTime = segment.ArrivalDateTime
			flight.DepartureTime = segment.DepartDateTime
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
				flight.DepartureTime = segment.DepartInfo.Time.DateTime
				flight.ArrivalTime = segment.ArrivalInfo.Time.DateTime
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
