package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type GetDataWithViaCepApiUseCaseImpl struct {
	Client *http.Client
}

func NewGetDataWithViaCepApiUseCaseImpl(client *http.Client) *GetDataWithViaCepApiUseCaseImpl {
	return &GetDataWithViaCepApiUseCaseImpl{
		Client: client,
	}
}

func (g *GetDataWithViaCepApiUseCaseImpl) Execute(ctx context.Context, zipCode string) (*ViaCep, *http.Response, error) {
	var response ViaCep

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", zipCode)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %v", err)
	}

	res, err := g.Client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, res, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response, res, nil
}
