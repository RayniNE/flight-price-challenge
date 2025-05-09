package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/raynine/flight-price-challenge/models"
)

func GetSkyScannerFlights() (*models.GetAutoCompleteResponse, error) {
	key := os.Getenv("SKY_SCANNER_KEY")
	host := os.Getenv("SKY_SCANNER_HOST")
	url := fmt.Sprintf("%s/flights/auto-complete?query=New", host)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	defer res.Body.Close()

	response := &models.GetAutoCompleteResponse{}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		return nil, err
	}

	return response, nil

}
