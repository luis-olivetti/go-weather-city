package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemperature_SetCelsius(t *testing.T) {
	// Arrange
	temp := &Temperature{}

	// Act
	temp.SetCelsius(25.0)

	// Assert
	assert.Equal(t, 25.0, temp.Celsius)
	assert.Equal(t, 77.0, temp.Fahrenheit)
	assert.Equal(t, 298.15, temp.Kelvin)
}
