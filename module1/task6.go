package main

import "fmt"

func main() {
	numbers := []int{42, -17, 3, 88, 0}
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	fmt.Println("Максимальное значение = ", max)
}
