package main

import "math"

// FuelForModule calculates the required fuel for a module given its mass.
func FuelForModule(mass uint64) uint64 {
	fuel := math.Floor(float64(mass)/3) - 2
	if fuel < 0 {
		return 0
	}
	return uint64(fuel)
}

// TotalFuelForModule calculates the required fuel for a module given its mass,
// including the weight of the fuel added.
func TotalFuelForModule(mass uint64) uint64 {
	var total uint64

	cur := mass

	for {
		fuel := FuelForModule(cur)
		if fuel == 0 {
			return total
		}

		total += fuel
		cur = fuel
	}
}
