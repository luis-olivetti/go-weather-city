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

type GetDataWithViaCepApiUseCase struct {
	Client *http.Client
}

func NewGetDataWithViaCepApiUseCase(client *http.Client) *GetDataWithViaCepApiUseCase {
	return &GetDataWithViaCepApiUseCase{
		Client: client,
	}
}

func (g *GetDataWithViaCepApiUseCase) Execute(ctx context.Context, zipCode string) (*ViaCep, error) {
	var response ViaCep

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", zipCode)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	res, err := g.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &response, nil
}
