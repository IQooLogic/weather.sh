package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Weather struct {
}

func GetWeather(lat, lng float32) (*WeatherInfo, error) {
	var wi *WeatherInfo

	url := fmt.Sprintf(API_URL, lat, lng, KEY)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Forbidden (most likely invalid token)")
	}

	err = json.NewDecoder(resp.Body).Decode(&wi)
	if err != nil {
		return nil, err
	}

	return wi, nil
}
