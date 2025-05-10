package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/raynine/flight-price-challenge/models"
)

func GetGoogleFlights(dto models.GetFlightsDTO) (*models.GoogleResponse, error) {
	key := os.Getenv("GOOGLE_FLIGHT_KEY")
	host := os.Getenv("GOOGLE_FLIGHT_HOST")

	url := fmt.Sprintf("https://%s/api/v1/searchFlights", host)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)

	query := req.URL.Query()

	query.Add("outbound_date", dto.DepartDate)
	query.Add("departure_id", dto.Origin)
	query.Add("arrival_id", dto.Destination)

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	defer res.Body.Close()

	response := &models.GoogleResponse{}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println("Error:", bodyString)

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println("Google API returned an error...")
		errResponse := &models.GoogleError{}

		err = json.NewDecoder(res.Body).Decode(&errResponse)
		if err != nil {
			log.Fatalf("Error occurred: %s", err.Error())
			return nil, err
		}

		fmt.Printf("An error ocurred while getting google flights: %s", response.Message)
		return nil, errors.New("error while retrieving flights")
	}

	return response, nil

}

func GetPricelineFlights(dto models.GetFlightsDTO) (*models.ResponsePriceline, error) {
	key := os.Getenv("PRICELINE_HOST")
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

	response := &models.ResponsePriceline{}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Printf("An error ocurred while decoding flights: %s", err.Error())
		return nil, errors.New("error while retrieving flights")
	}

	if !response.Status {
		fmt.Printf("An error ocurred while getting Priceline flights: %s", response.Errors)
		return nil, fmt.Errorf("an error ocurred while getting Priceline flights: %s", response.Errors)
	}

	return response, nil

}
