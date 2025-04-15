package main

import (
	"fmt"
	"math"
)

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {
	speedInMps := 120.4
	speedInKps := speedInMps * 3.6                                                      // преобразование
	var euroSpeed EuropeanVelocity = EuropeanVelocity(math.Round(speedInKps*100) / 100) // округление

	speedInMps2 := 130.0
	speedInKph := speedInMps2 * 2.23694                                                // преобразование
	var usaSpeed AmericanVelocity = AmericanVelocity(math.Round(speedInKph*100) / 100) // округление
	fmt.Println(usaSpeed)
	fmt.Println(euroSpeed)
}
