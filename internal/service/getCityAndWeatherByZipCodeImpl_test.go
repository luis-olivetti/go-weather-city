package service

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/luis-olivetti/go-weather-city/internal/entity"
	"github.com/luis-olivetti/go-weather-city/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGetDataWithViaCepApiUseCase struct {
	mock.Mock
}

func (m *MockGetDataWithViaCepApiUseCase) Execute(ctx context.Context, zipCode string) (*usecase.ViaCep, *http.Response, error) {
	args := m.Called(ctx, zipCode)
	return args.Get(0).(*usecase.ViaCep), args.Get(1).(*http.Response), args.Error(2)
}

type MockGetTemperatureWithWeatherApiUseCase struct {
	mock.Mock
}

func (m *MockGetTemperatureWithWeatherApiUseCase) Execute(ctx context.Context, cityName string) (*entity.Temperature, error) {
	args := m.Called(ctx, cityName)
	return args.Get(0).(*entity.Temperature), args.Error(1)
}

func TestGetCityAndWeatherByZipCodeDTO_Execute_ContainsLocalidade(t *testing.T) {
	// Arrange
	getDataWithViaCepMock := new(MockGetDataWithViaCepApiUseCase)
	getTemperatureWithWeatherMock := new(MockGetTemperatureWithWeatherApiUseCase)

	viaCepFake := &usecase.ViaCep{
		Cep:        "12345678",
		Logradouro: "Rua Teste",
		Bairro:     "Bairro Teste",
		Localidade: "Teste",
		Uf:         "SP",
	}

	temperatureFake := &entity.Temperature{
		Celsius:    10,
		Fahrenheit: 50,
		Kelvin:     283.15,
	}

	getDataWithViaCepMock.On("Execute", context.Background(), "12345678").Return(viaCepFake, &http.Response{}, nil)
	getTemperatureWithWeatherMock.On("Execute", context.Background(), "Teste").Return(temperatureFake, nil)

	usecase := NewGetCityAndWeatherByZipCodeImpl(getDataWithViaCepMock, getTemperatureWithWeatherMock)

	// Act
	result, _, statusCode := usecase.Execute(context.Background(), "12345678")

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "12345678", result.ZipCode)
	assert.Equal(t, "Teste", result.CityName)
	assert.Equal(t, 10.0, result.CelsiusTemperature)
	assert.Equal(t, 50.0, result.FahrenheitTemperature)
	assert.Equal(t, 283.15, result.KelvinTemperature)
	assert.Equal(t, int16(200), statusCode)
}

func TestGetCityAndWeatherByZipCodeDTO_Execute_NotContainsLocalidade(t *testing.T) {
	// Arrange
	getDataWithViaCepMock := new(MockGetDataWithViaCepApiUseCase)
	getTemperatureWithWeatherMock := new(MockGetTemperatureWithWeatherApiUseCase)

	viaCepFake := &usecase.ViaCep{}

	temperatureFake := &entity.Temperature{
		Celsius:    10,
		Fahrenheit: 50,
		Kelvin:     283.15,
	}

	getDataWithViaCepMock.On("Execute", context.Background(), "12345678").Return(viaCepFake, &http.Response{}, nil)
	getTemperatureWithWeatherMock.On("Execute", context.Background(), "Teste").Return(temperatureFake, nil)

	usecase := NewGetCityAndWeatherByZipCodeImpl(getDataWithViaCepMock, getTemperatureWithWeatherMock)

	// Act
	result, _, statusCode := usecase.Execute(context.Background(), "12345678")

	// Assert
	assert.Nil(t, result)
	assert.Equal(t, int16(404), statusCode)
}

func TestGetCityAndWeatherByZipCodeDTO_Execute_ViaCepEmitsHttpCode422(t *testing.T) {
	// Arrange
	getDataWithViaCepMock := new(MockGetDataWithViaCepApiUseCase)
	getTemperatureWithWeatherMock := new(MockGetTemperatureWithWeatherApiUseCase)

	viaCepFake := &usecase.ViaCep{}

	responseFake := &http.Response{
		StatusCode: 422,
	}

	getDataWithViaCepMock.On("Execute", context.Background(), "123456789").Return(viaCepFake, responseFake, errors.New("Invalid ZipCode"))

	usecase := NewGetCityAndWeatherByZipCodeImpl(getDataWithViaCepMock, getTemperatureWithWeatherMock)

	// Act
	result, err, statusCode := usecase.Execute(context.Background(), "123456789")

	// Assert
	assert.Nil(t, result)
	assert.Equal(t, "Invalid ZipCode", err.Error())
	assert.Equal(t, int16(422), statusCode)
}

func TestGetCityAndWeatherByZipCodeDTO_Execute_ViaCepEmitsHttpCode400(t *testing.T) {
	// Arrange
	getDataWithViaCepMock := new(MockGetDataWithViaCepApiUseCase)
	getTemperatureWithWeatherMock := new(MockGetTemperatureWithWeatherApiUseCase)

	viaCepFake := &usecase.ViaCep{}

	responseFake := &http.Response{
		StatusCode: 400,
	}

	getDataWithViaCepMock.On("Execute", context.Background(), "x").Return(viaCepFake, responseFake, errors.New("Bad Request"))

	usecase := NewGetCityAndWeatherByZipCodeImpl(getDataWithViaCepMock, getTemperatureWithWeatherMock)

	// Act
	result, err, statusCode := usecase.Execute(context.Background(), "x")

	// Assert
	assert.Nil(t, result)
	assert.Equal(t, "ZipCode not found", err.Error())
	assert.Equal(t, int16(404), statusCode)
}

func TestGetCityAndWeatherByZipCodeDTO_Execute_ViaCepEmitsHttpCode500(t *testing.T) {
	// Arrange
	getDataWithViaCepMock := new(MockGetDataWithViaCepApiUseCase)
	getTemperatureWithWeatherMock := new(MockGetTemperatureWithWeatherApiUseCase)

	viaCepFake := &usecase.ViaCep{}

	responseFake := &http.Response{
		StatusCode: 500,
	}

	getDataWithViaCepMock.On("Execute", context.Background(), "12345678").Return(viaCepFake, responseFake, errors.New("Internal Server Error"))

	usecase := NewGetCityAndWeatherByZipCodeImpl(getDataWithViaCepMock, getTemperatureWithWeatherMock)

	// Act
	result, _, statusCode := usecase.Execute(context.Background(), "12345678")

	// Assert
	assert.Nil(t, result)
	assert.Equal(t, int16(500), statusCode)
}

func TestGetCityAndWeatherByZipCodeDTO_Execute_WeatherEmitsError(t *testing.T) {
	// Arrange
	getDataWithViaCepMock := new(MockGetDataWithViaCepApiUseCase)
	getTemperatureWithWeatherMock := new(MockGetTemperatureWithWeatherApiUseCase)

	viaCepFake := &usecase.ViaCep{
		Cep:        "12345678",
		Logradouro: "Rua Teste",
		Bairro:     "Bairro Teste",
		Localidade: "Teste",
		Uf:         "SP",
	}

	temperatureFake := &entity.Temperature{}

	getDataWithViaCepMock.On("Execute", context.Background(), "12345678").Return(viaCepFake, &http.Response{}, nil)
	getTemperatureWithWeatherMock.On("Execute", context.Background(), "Teste").Return(temperatureFake, errors.New("Internal Server Error"))

	usecase := NewGetCityAndWeatherByZipCodeImpl(getDataWithViaCepMock, getTemperatureWithWeatherMock)

	// Act
	result, _, statusCode := usecase.Execute(context.Background(), "12345678")

	// Assert
	assert.Nil(t, result)
	assert.Equal(t, int16(500), statusCode)
}
