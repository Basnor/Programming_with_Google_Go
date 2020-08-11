package main

import (
	"fmt"
	"math"
)

func main() {
	accel, initVel, initDisp := getParamsFromConsole()
	displacement := GenDisplaceFn(accel, initVel, initDisp)

	for {
		var fTime float64
		fmt.Print("Time: ")
		fmt.Scanln(&fTime)
		fmt.Println(displacement(fTime))
	}
}

func GenDisplaceFn(acc, initV, initD float64) func(t float64) float64 {
	rFunc := func(time float64) float64 {
		return 0.5 * acc * math.Pow(time, 2.0) + initV * time + initD
	}
	return rFunc
}

func getParamsFromConsole() (float64, float64, float64) {
	NameVal := map[string]float64 {
		"Acceleration":0.0,
		"Initial Velocity":0.0,
		"Initial Displacement":0.0,
	}

	for name, val := range NameVal {
		fmt.Printf("%s: ", name)
		fmt.Scanln(&val)
		NameVal[name] = val
	}

	return NameVal["Acceleration"], NameVal["Initial Velocity"], NameVal["Initial Displacement"]
}