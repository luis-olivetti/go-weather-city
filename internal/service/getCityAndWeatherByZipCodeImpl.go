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

func (g *GetCityAndWeatherByZipCodeImpl) Execute(ctx context.Context, zipCode string) *GetCityAndWeatherByZipCodeDTO {
	result, err := g.GetDataWithViaCepApiUseCase.Execute(ctx, zipCode)
	if err != nil {
		fmt.Println(err)
	}

	temperature, err := g.GetTemperatureWithWeatherApiUseCase.Execute(ctx, result.Localidade)
	if err != nil {
		fmt.Println(err)
	}

	dto := &GetCityAndWeatherByZipCodeDTO{
		ZipCode:               result.Cep,
		CityName:              result.Localidade,
		CelsiusTemperature:    temperature.Celsius,
		FahrenheitTemperature: temperature.Fahrenheit,
		KelvinTemperature:     temperature.Kelvin,
	}

	return dto
}
