package service

import (
	"context"
	"fmt"

	"github.com/luis-olivetti/go-weather-city/internal/usecase"
)

type GetCityAndWeatherByZipCodeDTO struct {
	ZipCode               string
	CityName              string
	CelsiusTemperature    float64 `json:"temp_C"`
	FahrenheitTemperature float64 `json:"temp_F"`
	KelvinTemperature     float64 `json:"temp_K"`
}

type GetCityAndWeatherByZipCodeImpl struct {
	GetDataWithViaCepApiUseCase         usecase.GetDataWithViaCepApiUseCaseInterface
	GetTemperatureWithWeatherApiUseCase usecase.GetTemperatureWithWeatherApiUseCaseInterface
}

func NewGetCityAndWeatherByZipCodeImpl(
	viaCepUseCase usecase.GetDataWithViaCepApiUseCaseInterface,
	weatherUseCase usecase.GetTemperatureWithWeatherApiUseCaseInterface,
) *GetCityAndWeatherByZipCodeImpl {
	return &GetCityAndWeatherByZipCodeImpl{
		GetDataWithViaCepApiUseCase:         viaCepUseCase,
		GetTemperatureWithWeatherApiUseCase: weatherUseCase,
	}
}

func (g *GetCityAndWeatherByZipCodeImpl) Execute(ctx context.Context, zipCode string) (*GetCityAndWeatherByZipCodeDTO, error, int16) {
	viaCep, res, err := g.GetDataWithViaCepApiUseCase.Execute(ctx, zipCode)
	if err != nil {
		if res.StatusCode >= 400 && res.StatusCode < 500 {

			if res.StatusCode == 422 {
				return nil, fmt.Errorf("Invalid ZipCode"), 422
			}

			return nil, fmt.Errorf("ZipCode not found"), 404
		}

		return nil, err, 500
	}

	if viaCep.Localidade == "" {
		return nil, fmt.Errorf("ZipCode not found"), 404
	}

	temperature, err := g.GetTemperatureWithWeatherApiUseCase.Execute(ctx, viaCep.Localidade)
	if err != nil {
		fmt.Println(err)
		return nil, err, 500
	}

	dto := &GetCityAndWeatherByZipCodeDTO{
		ZipCode:               viaCep.Cep,
		CityName:              viaCep.Localidade,
		CelsiusTemperature:    temperature.Celsius,
		FahrenheitTemperature: temperature.Fahrenheit,
		KelvinTemperature:     temperature.Kelvin,
	}

	return dto, nil, 200
}
