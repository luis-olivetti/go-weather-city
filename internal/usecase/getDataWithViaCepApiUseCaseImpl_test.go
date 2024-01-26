package usecase

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetDataWithViaCepApiUseCaseImpl_Execute_ValidZipCode(t *testing.T) {
	// Arrange
	mockResponseBody := `{"cep":"12345678","logradouro":"Rua Teste","bairro":"Bairro Teste","localidade":"Cidade Teste","uf":"TS"}`
	mockResponse := httpmock.NewStringResponse(http.StatusOK, mockResponseBody)

	httpmock.RegisterResponder("GET", "http://viacep.com.br/ws/12345678/json/", httpmock.ResponderFromResponse(mockResponse))

	client := &http.Client{Transport: httpmock.DefaultTransport}
	useCase := NewGetDataWithViaCepApiUseCaseImpl(client)

	// Act
	result, _, err := useCase.Execute(context.Background(), "12345678")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "12345678", result.Cep)
	assert.Equal(t, "Rua Teste", result.Logradouro)
	assert.Equal(t, "Bairro Teste", result.Bairro)
	assert.Equal(t, "Cidade Teste", result.Localidade)
	assert.Equal(t, "TS", result.Uf)
}

func TestGetDataWithViaCepApiUseCaseImpl_Execute_InvalidZipCode(t *testing.T) {
	// Arrange
	client := &http.Client{Transport: httpmock.DefaultTransport}
	useCase := NewGetDataWithViaCepApiUseCaseImpl(client)

	// Act
	result, res, err := useCase.Execute(context.Background(), "123456789")

	// Assert
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	assert.Error(t, err)
	assert.Nil(t, result)
}
