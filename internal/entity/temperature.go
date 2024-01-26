package entity

type Temperature struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

func (t *Temperature) SetCelsius(temp float64) {
	t.Celsius = temp
	t.calcFahrenheit()
	t.calcKelvin()
}

func (t *Temperature) calcFahrenheit() {
	t.Fahrenheit = t.Celsius*1.8 + 32
}

func (t *Temperature) calcKelvin() {
	t.Kelvin = t.Celsius + 273.15
}
