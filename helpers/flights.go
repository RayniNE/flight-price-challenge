package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/raynine/flight-price-challenge/models"
)

func GetPricelineFlights(dto models.GetFlightsDTO) ([]models.Flights, error) {
	key := os.Getenv("RAPID_API_KEY")
	host := os.Getenv("PRICELINE_KEY")

	url := fmt.Sprintf("https://%s/flights/search-one-way", host)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)

	query := req.URL.Query()

	query.Add("departureDate", dto.DepartDate)
	query.Add("originAirportCode", dto.Origin)
	query.Add("destinationAirportCode", dto.Destination)

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	defer res.Body.Close()

	model := &models.PricelineResponse{}

	err = json.NewDecoder(res.Body).Decode(&model)
	if err != nil {
		fmt.Printf("An error ocurred while decoding flights: %s", err.Error())
		return nil, errors.New("error while retrieving flights")
	}

	response := model.MapPriceLineToModel()

	return response, nil
}

func GetFlightSkyFlights(dto models.GetFlightsDTO) ([]models.Flights, error) {
	key := os.Getenv("RAPID_API_KEY")
	host := os.Getenv("FLIGHTS_SKY_KEY")

	url := fmt.Sprintf("https://%s/flights/search-one-way", host)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)

	query := req.URL.Query()

	query.Add("departDate", dto.DepartDate)
	query.Add("fromEntityId", dto.Origin)
	query.Add("toEntityId", dto.Destination)

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	defer res.Body.Close()

	model := &models.ResponseFlightsSky{}

	err = json.NewDecoder(res.Body).Decode(&model)
	if err != nil {
		fmt.Printf("An error ocurred while decoding flights: %s", err.Error())
		return nil, errors.New("error while retrieving flights")
	}

	response := model.MapPriceLineToModel()

	return response, nil
}

func GetAgodaFlights(dto models.GetFlightsDTO) ([]models.Flights, error) {
	key := os.Getenv("RAPID_API_KEY")
	host := os.Getenv("AGODA_KEY")

	url := fmt.Sprintf("https://%s/flights/search-one-way", host)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)

	query := req.URL.Query()

	query.Add("departureDate", dto.DepartDate)
	query.Add("origin", dto.Origin)
	query.Add("destination", dto.Destination)

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	defer res.Body.Close()

	model := &models.AgodaResponse{}

	err = json.NewDecoder(res.Body).Decode(&model)
	if err != nil {
		fmt.Printf("An error ocurred while decoding flights: %s", err.Error())
		return nil, errors.New("error while retrieving flights")
	}

	response := model.MapPriceLineToModel()

	return response, nil
}

func GetFlightsResponse(dto models.GetFlightsDTO) (*models.FlightsResponse, error) {
	model := &models.FlightsByAPI{}
	var (
		priceLineError    error
		agodaFlightsError error
		flightsSkyError   error
	)

	wg := sync.WaitGroup{}

	wg.Add(3)

	go func(model *models.FlightsByAPI) {
		defer wg.Done()
		priceLineResponse, err := GetPricelineFlights(dto)
		if err != nil {
			priceLineError = fmt.Errorf("an error ocurred while getting priceline flights: %s", err.Error())
		}

		model.PriceLine = priceLineResponse
	}(model)

	go func(model *models.FlightsByAPI) {
		defer wg.Done()

		flightsSkyResponse, err := GetFlightSkyFlights(dto)
		if err != nil {
			flightsSkyError = fmt.Errorf("an error ocurred while getting flights sky flights: %s", err.Error())
		}

		model.FlightSky = flightsSkyResponse
	}(model)

	go func(model *models.FlightsByAPI) {
		defer wg.Done()
		agodaResponse, err := GetAgodaFlights(dto)
		if err != nil {
			agodaFlightsError = fmt.Errorf("an error ocurred while getting agoda flights: %s", err.Error())
		}

		model.Agoda = agodaResponse
	}(model)

	wg.Wait()

	if agodaFlightsError != nil {
		return nil, agodaFlightsError
	}

	if priceLineError != nil {
		return nil, priceLineError
	}

	if flightsSkyError != nil {
		return nil, flightsSkyError
	}

	priceLinePriceOrdered := GetOrderedFlightByPrice(model.PriceLine)
	agodaPriceOrdered := GetOrderedFlightByPrice(model.Agoda)
	flightsSkyPriceOrdered := GetOrderedFlightByPrice(model.FlightSky)

	cheapestFlights := []models.Flights{priceLinePriceOrdered[0], agodaPriceOrdered[0], flightsSkyPriceOrdered[0]}

	priceLineArrivalOrdered := GetOrderedFlightByTime(priceLinePriceOrdered)
	agodaArrivalOrdered := GetOrderedFlightByTime(agodaPriceOrdered)
	flightsSkyArrivalOrdered := GetOrderedFlightByTime(flightsSkyPriceOrdered)

	fastestFlights := []models.Flights{priceLineArrivalOrdered[0], agodaArrivalOrdered[0], flightsSkyArrivalOrdered[0]}

	// Sort the final slices by price and fastest
	cheapestFlights = GetOrderedFlightByPrice(cheapestFlights)

	fastestFlights = GetOrderedFlightByTime(fastestFlights)

	otherFlights := []models.Flights{}
	otherFlights = append(otherFlights, flightsSkyPriceOrdered[1:]...)
	otherFlights = append(otherFlights, agodaPriceOrdered[1:]...)
	otherFlights = append(otherFlights, flightsSkyPriceOrdered[1:]...)
	otherFlights = append(otherFlights, priceLineArrivalOrdered[1:]...)
	otherFlights = append(otherFlights, agodaArrivalOrdered[1:]...)
	otherFlights = append(otherFlights, flightsSkyArrivalOrdered[1:]...)

	response := &models.FlightsResponse{
		CheapestFlights: cheapestFlights,
		FastestFlights:  fastestFlights,
		OtherFlights:    otherFlights,
	}

	return response, nil
}
