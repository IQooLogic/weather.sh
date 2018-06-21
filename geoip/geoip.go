package geoip

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type geoip struct {
}

func GetLocation() (location Location, err error) {
	resp, err := http.Get(API_URL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return
	}

	if location.Status != "success" {
		return location, fmt.Errorf("\nunable to determine location")
	}

	return location, nil
}
