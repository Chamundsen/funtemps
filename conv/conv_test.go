package conv

import (
	"math"
	"testing"
)

/*
*

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
		{input: 15, want: -9.44},
		{input: 100, want: 37.78},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// De andre testfunksjonene implementeres her

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 134.01},
		{input: 23, want: 73.4},
		{input: 10, want: 50.0},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 26.85, want: 300.00},
		{input: 10.00, want: 283.15},
		{input: 20.00, want: 293.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 300.00, want: 26.85},
		{input: 320.00, want: 46.85},
		{input: 240.00, want: -33.14},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 50, want: 283.15},
		{input: 60, want: 288.70},
		{input: 70, want: 294.26},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 283.15, want: 50.00},
		{input: 270.15, want: 26.60},
		{input: 260.15, want: 8.60},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
