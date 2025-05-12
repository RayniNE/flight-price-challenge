package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

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

	response := GetOrderedPriceLineFlightsByPrice(model)

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

	response := GetOrderedFlightsSkyByPrice(model)

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

	response := GetOrderedAgodaFlightsByPrice(model)

	return response, nil
}
