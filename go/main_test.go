package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestFuelCalculator(t *testing.T) {
	is := is.New(t)

	tests := map[float64]float64{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}

	for k, v := range tests {
		t.Logf("testing mass %f requires fuel %f", k, v)
		is.Equal(fuelCalculator(k), v)
		t.Log("ok")
	}
}

func TestTotalFuelCalculator(t *testing.T) {
	is := is.New(t)

	tests := map[float64]float64{
		14:     2,
		1969:   966,
		100756: 50346,
	}

	for k, v := range tests {
		t.Logf("testing mass %f requires fuel %f", k, v)
		is.Equal(totalFuelCalculator(k), v)
		t.Log("ok")
	}
}
