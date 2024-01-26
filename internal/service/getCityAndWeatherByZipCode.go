package service

import (
	"context"
	"fmt"

	"github.com/luis-olivetti/go-weather-city/internal/usecase"
)

type GetCityAndWeatherByZipCode struct{}

func NewGetCityAndWeatherByZipCode() *GetCityAndWeatherByZipCode {
	return &GetCityAndWeatherByZipCode{}
}

func (g *GetCityAndWeatherByZipCode) Execute(ctx context.Context, zipCode string) string {
	usecase := &usecase.GetDataWithViaCepApiUseCase{}

	resultado, err := usecase.Execute(ctx, zipCode)
	if err != nil {
		fmt.Println(err)
	}
	return resultado.Logradouro
}
