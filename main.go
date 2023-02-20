package main

import (
	"flag"
	"fmt"

	"github.com/Chamundsen/funtemps/conv"
)

var fahr float64
var cel float64
var kel float64
var out string
var funfacts string

func printTemperatures() {
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

func init() {
	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	// Du må selv definere flag-variablene for "C" og "K".
	flag.StringVar(&out, "out", "0.0", "beregne temperatur i C - Celsius, F - Fahrenheit, K - Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"funfacts\" om sun - solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t-flagget, som bestemmer
	// hvilken temperaturskala du skal bruke når du viser funfacts

}

func main() {
	flag.Parse()
	printTemperatures()

	if isFlagPassed("F") && out == "C" {
		celsius := conv.FarhenheitToCelsius(fahr)
		fmt.Printf("F: %.2f til C: %.2f\n", fahr, celsius)
	} else if isFlagPassed("F") && out == "K" {
		kelvin := conv.FahrenheitToKelvin(fahr)
		fmt.Printf("F: %.2f til K: %.2f\n", fahr, kelvin)
	} else if isFlagPassed("C") && out == "F" {
		fahrenheit := conv.CelsiusToFahrenheit(cel)
		fmt.Printf("C: %.2f til F: %.2f\n", cel, fahrenheit)
	} else if isFlagPassed("C") && out == "K" {
		kelvin := conv.CelsiusToFahrenheit(cel)
		fmt.Printf("C: %.2f til K: %.2f\n", cel, kelvin)
	} else if isFlagPassed("K") && out == "C" {
		celsius := conv.CelsiusToFahrenheit(kel)
		fmt.Printf("K: %.2f til C: %.2f\n", kel, celsius)
	} else if isFlagPassed("K") && out == "F" {
		fahrenheit := conv.CelsiusToFahrenheit(kel)
		fmt.Printf("K: %.2f til F: %.2f\n", kel, fahrenheit)
	} else {
		fmt.Println("Noe gikk galt")
	}

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
