package utils

type (
	Celsius    float64
	Kelvin     float64
	Fahrenheit float64
)

func Kelvin2Celsius(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
