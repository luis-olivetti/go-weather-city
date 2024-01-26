package usecase

import (
	"context"
	"net/http"
)

type GetDataWithViaCepApiUseCaseInterface interface {
	Execute(ctx context.Context, zipCode string) (*ViaCep, *http.Response, error)
}
