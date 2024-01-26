package usecase

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetTemperatureWithWeatherApiUseCaseImpl_Execute(t *testing.T) {
	// Arrange
	mockResponseBody := `{
		"location": {
		  "name": "Example City",
		  "region": "Example Region",
		  "country": "Example Country",
		  "lat": 123.456,
		  "lon": -78.901,
		  "tz_id": "Example Timezone",
		  "localtime_epoch": 1616700000,
		  "localtime": "2022-03-25 12:00:00"
		},
		"current": {
		  "temp_c": 25.5,
		  "condition": {}
		}
	  }`
	mockResponse := httpmock.NewStringResponse(http.StatusOK, mockResponseBody)

	httpmock.RegisterResponder("GET", "http://api.weatherapi.com/v1/current.json", httpmock.ResponderFromResponse(mockResponse))

	client := &http.Client{Transport: httpmock.DefaultTransport}
	useCase := NewGetTemperatureWithWeatherApiUseCaseImpl(client)

	// Act
	result, err := useCase.Execute(context.Background(), "Example City")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 25.5, result.Celsius)
	assert.Equal(t, 77.9, result.Fahrenheit)
	assert.Equal(t, 298.65, result.Kelvin)
}

func TestGetTemperatureWithWeatherApiUseCaseImpl_Execute_CityNameNotFound(t *testing.T) {
	// Arrange
	mockResponseBody := `{
		"error": {
		  "code": 1006,
		  "message": "No matching location found."
		}
	  }`
	mockResponse := httpmock.NewStringResponse(http.StatusNotFound, mockResponseBody)

	httpmock.RegisterResponder("GET", "http://api.weatherapi.com/v1/current.json", httpmock.ResponderFromResponse(mockResponse))

	client := &http.Client{Transport: httpmock.DefaultTransport}
	useCase := NewGetTemperatureWithWeatherApiUseCaseImpl(client)

	// Act
	result, err := useCase.Execute(context.Background(), "City Name Not Found")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}
