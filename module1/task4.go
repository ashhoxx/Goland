package main

import "fmt"

func main() {
	var a int
	b := 1
	fmt.Print("От какого числа вы хотите найти факториал? ")
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		b *= i
	}
	fmt.Println("Факториал от числа ", a, "равен = ", b)
}
