package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende

	celsius := (value - 32) * (5.00 / 9.00)
	return celsius
}

// De andre konverteringsfunksjonene implementere her
func CelsiusToFahrenheit(value float64) float64 {
	fahrenheit := (value * 9.00 / 5.00) + 32
	return fahrenheit
}

func CelsiusToKelvin(value float64) float64 {
	kelvin := (value + 273.15)
	return kelvin
}

func KelvinToCelsius(value float64) float64 {
	celsius1 := (value - 273.15)
	return celsius1
}

func FahrenheitToKelvin(value float64) float64 {
	kelvin1 := (value-32)*(5.00/9.00) + 273.15
	return kelvin1
}

func KelvinToFahrenheit(value float64) float64 {
	fahrenheit1 := (value-273.15)*(9.00/5.00) + 32
	return fahrenheit1
}
