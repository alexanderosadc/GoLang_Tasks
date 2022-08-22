package main

import (
	"fmt"
	"math"
)

func main() {
	var a, v0, s0, t float64
	UserInput(&a, &v0, &s0, &t)

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("Displacement = ", fn(t))
}

func UserInput(accelearation, velocity, displacement, time *float64) {
	fmt.Println("Write value for acceleration")
	fmt.Scan(accelearation)
	fmt.Println("Write value for velocity")
	fmt.Scan(velocity)
	fmt.Println("Write value for displacement")
	fmt.Scan(displacement)
	fmt.Println("Write value for time which will calculate after specific time displacement")
	fmt.Scan(time)
}

// Function which has return value another function with one parameter time
func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (v0 * t) + s0
	}
}
