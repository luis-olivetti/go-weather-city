package usecase

import (
	"context"

	"github.com/luis-olivetti/go-weather-city/internal/entity"
)

type GetTemperatureWithWeatherApiUseCaseInterface interface {
	Execute(ctx context.Context, cityName string) (*entity.Temperature, error)
}
