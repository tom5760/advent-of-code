package main

import (
	"fmt"
	"testing"
)

func TestFuelForModule(t *testing.T) {
	tests := []struct {
		Mass uint64
		Fuel uint64
	}{
		{Mass: 12, Fuel: 2},
		{Mass: 14, Fuel: 2},
		{Mass: 1969, Fuel: 654},
		{Mass: 100756, Fuel: 33583},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := FuelForModule(tt.Mass)

			if tt.Fuel != actual {
				t.Errorf("expected %v != actual %v", tt.Fuel, actual)
			}
		})
	}
}

func TestTotalFuelForModule(t *testing.T) {
	tests := []struct {
		Mass uint64
		Fuel uint64
	}{
		{Mass: 14, Fuel: 2},
		{Mass: 1969, Fuel: 966},
		{Mass: 100756, Fuel: 50346},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			actual := TotalFuelForModule(tt.Mass)

			if tt.Fuel != actual {
				t.Errorf("expected %v != actual %v", tt.Fuel, actual)
			}
		})
	}
}
