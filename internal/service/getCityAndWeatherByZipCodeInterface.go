package service

import (
	"context"
)

type GetCityAndWeatherByZipCodeInteface interface {
	Execute(ctx context.Context, zipCode string) string
}
