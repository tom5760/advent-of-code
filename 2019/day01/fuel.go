package main

// FuelForModule calculates the required fuel for a module given its mass.
//
// Fuel required to launch a given module is based on its mass. Specifically,
// to find the fuel required for a module, take its mass, divide by three,
// round down, and subtract 2.
func FuelForModule(mass uint64) uint64 {
	// Go integer division truncates towards zero.
	fuel := mass / 3

	if fuel < 2 {
		return 0
	}

	return fuel - 2
}

// TotalFuelForModule calculates the required fuel for a module given its mass,
// including the weight of the fuel added.
func TotalFuelForModule(mass uint64) uint64 {
	var total uint64

	for {
		mass = FuelForModule(mass)
		if mass == 0 {
			return total
		}

		total += mass
	}
}
