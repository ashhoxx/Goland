package main

import "fmt"

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(getMax(1, 5))
}
