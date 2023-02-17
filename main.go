package main

import (
	"fmt"

	"github.com/Chamundsen/funtemps/conv"
)

var fahrenheit float64
var out string
var funfacts string

func main() {
	// ...
	fahrenheitFloat := 134.0
	celsius := conv.FarhenheitToCelsius(fahrenheitFloat)
	fmt.Printf("%.2f °C\n", celsius)

	celsiusFloat := 56.67
	fahrenheit := conv.CelsiusToFahrenheit(celsiusFloat)
	fmt.Printf("%.2f °F\n", fahrenheit)

	kelvinFloat := 26.85
	kelvin := conv.CelsiusToKelvin(kelvinFloat)
	fmt.Printf("%.2f °K\n", kelvin)

	celsiusFloat1 := 300.0
	celsius1 := conv.KelvinToCelsius(celsiusFloat1)
	fmt.Printf("%.2f °C\n", celsius1)

	kelvinFloat1 := 50.0
	kelvin1 := conv.FahrenheitToKelvin(kelvinFloat1)
	fmt.Printf("%.2f °K\n", kelvin1)

	fahrenheitFloat1 := 283.15
	fahrenheit1 := conv.KelvinToFahrenheit(fahrenheitFloat1)
	fmt.Printf("%.2f °F\n", fahrenheit1)

}
