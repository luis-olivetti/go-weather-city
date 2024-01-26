package usecase

import (
	"context"
)

type GetDataWithViaCepApiUseCaseInterface interface {
	Execute(ctx context.Context, zipCode string) (*ViaCep, error)
}
