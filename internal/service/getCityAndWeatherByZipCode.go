package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/luis-olivetti/go-weather-city/internal/usecase"
)

type GetCityAndWeatherByZipCode struct {
	Client *http.Client
}

func NewGetCityAndWeatherByZipCode(client *http.Client) *GetCityAndWeatherByZipCode {
	return &GetCityAndWeatherByZipCode{
		Client: client,
	}
}

func (g *GetCityAndWeatherByZipCode) Execute(ctx context.Context, zipCode string) string {
	usecase := usecase.NewGetDataWithViaCepApiUseCase(g.Client)

	resultado, err := usecase.Execute(ctx, zipCode)
	if err != nil {
		fmt.Println(err)
	}
	return resultado.Logradouro
}
