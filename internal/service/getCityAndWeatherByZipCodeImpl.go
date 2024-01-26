package service

import (
	"context"
	"fmt"

	"github.com/luis-olivetti/go-weather-city/internal/usecase"
)

type GetCityAndWeatherByZipCodeImpl struct {
	GetDataWithViaCepApiUseCase usecase.GetDataWithViaCepApiUseCaseInterface
}

func NewGetCityAndWeatherByZipCodeImpl(usecase usecase.GetDataWithViaCepApiUseCaseInterface) *GetCityAndWeatherByZipCodeImpl {
	return &GetCityAndWeatherByZipCodeImpl{
		GetDataWithViaCepApiUseCase: usecase,
	}
}

func (g *GetCityAndWeatherByZipCodeImpl) Execute(ctx context.Context, zipCode string) string {
	result, err := g.GetDataWithViaCepApiUseCase.Execute(ctx, zipCode)
	if err != nil {
		fmt.Println(err)
	}
	return result.Logradouro
}
