package main

import "fmt"

func contains(a []string, x string) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	return false
}
func main() {
	names := []string{"Bob", "Alice", "Kevin"}
	fmt.Println(contains(names, "Bob"))
	fmt.Println(contains(names, "Cop"))
}
