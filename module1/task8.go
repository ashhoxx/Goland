package main

import "fmt"

func main() {
	num := []int{44, 22, 55, 66, 77}
	sum := 0
	for _, i := range num {
		sum += i
	}
	fmt.Print("Сумма всех элементов массива = ", sum)
}
