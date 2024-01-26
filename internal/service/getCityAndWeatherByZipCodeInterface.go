package service

import (
	"context"
)

type GetCityAndWeatherByZipCodeInteface interface {
	Execute(ctx context.Context, zipCode string) (*GetCityAndWeatherByZipCodeDTO, error, int16)
}
