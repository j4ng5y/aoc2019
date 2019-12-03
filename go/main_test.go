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
		is.Equal(v, fuelCalculator(k))
	}
}
