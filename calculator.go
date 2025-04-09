package main

import "fmt"

func main() {
	var first int
	var second int
	var i string
	fmt.Print("Введите два числа = ")
	fmt.Scan(&first, &second)
	fmt.Print("Какую операцию вы хотите выполнить = ")
	fmt.Scan(&i)
	if i == "+" {
		fmt.Print("Сумма ваших чисел = ", first+second)
	}
	if i == "-" {
		fmt.Print("Разность ваших чисел = ", first-second)
	}
	if i == "*" {
		fmt.Print("Умножение ваших чисел = ", first*second)
	}
	if i == "/" {
		fmt.Print("Деление ваших чисел = ", first/second)
	}
}
