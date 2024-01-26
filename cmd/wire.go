//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/google/wire"
	"github.com/luis-olivetti/go-weather-city/internal/service"
	"github.com/luis-olivetti/go-weather-city/internal/usecase"
)

func InitializeGetDataWithViaCepApiUseCase(*http.Client) *usecase.GetDataWithViaCepApiUseCaseImpl {
	wire.Build(usecase.NewGetDataWithViaCepApiUseCaseImpl)
	return &usecase.GetDataWithViaCepApiUseCaseImpl{}
}

func InitializeGetTemperatureWithWeatherApiUseCase(*http.Client) *usecase.GetTemperatureWithWeatherApiUseCaseImpl {
	wire.Build(usecase.NewGetTemperatureWithWeatherApiUseCaseImpl)
	return &usecase.GetTemperatureWithWeatherApiUseCaseImpl{}
}

func InitializeGetCityAndWeatherByZipCode(
	viaCepUseCase usecase.GetDataWithViaCepApiUseCaseInterface,
	weatherUseCase usecase.GetTemperatureWithWeatherApiUseCaseInterface,
) *service.GetCityAndWeatherByZipCodeImpl {
	wire.Build(service.NewGetCityAndWeatherByZipCodeImpl)
	return &service.GetCityAndWeatherByZipCodeImpl{}
}
