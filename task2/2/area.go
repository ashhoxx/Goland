package main

import "fmt"

func main() {
	circumference := 35.0
	R := new(float64)

	*R = circumference * 3.14

	area := 3.14 * (*R) * (*R)

	fmt.Println(area)
	fmt.Println(*R)

}
