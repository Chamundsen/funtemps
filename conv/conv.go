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
	return (value - 32) * (5.0 / 9.0)
}

// De andre konverteringsfunksjonene implementere her
func CelsiusToFahrenheit(value float64) float64 {
	return (value * 9.0 / 5.0) + 32
}

func CelsiustoKelvin(value float64) float64 {
	return (value + 273.15)
}

func KelvintoCelsius(value float64) float64 {
	return (value - 273.15)
}

func FahrenheittoKelvin(value float64) float64 {
	return (value-32)*(5.0/9.0) + 273.15
}

func KelvintoFahrenheit(value float64) float64 {
	return (value-273.15)*(9.0/5.0) + 32
}
