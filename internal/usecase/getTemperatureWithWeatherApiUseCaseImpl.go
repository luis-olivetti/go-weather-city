package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luis-olivetti/go-weather-city/internal/entity"
)

type Weather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
		} `json:"condition"`
	} `json:"current"`
}

type GetTemperatureWithWeatherApiUseCaseImpl struct {
	Client *http.Client
}

func NewGetTemperatureWithWeatherApiUseCaseImpl(client *http.Client) *GetTemperatureWithWeatherApiUseCaseImpl {
	return &GetTemperatureWithWeatherApiUseCaseImpl{
		Client: client,
	}
}

func (g *GetTemperatureWithWeatherApiUseCaseImpl) Execute(ctx context.Context, cityName string) (*entity.Temperature, error) {
	var response Weather

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=a91eb948a337442782b123810242601&q=%s", cityName)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	res, err := g.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	temperature := &entity.Temperature{}
	temperature.SetCelsius(response.Current.TempC)

	return temperature, nil
}
