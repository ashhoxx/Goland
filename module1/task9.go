package main

import "fmt"

func main() {
	fib := []int{0, 1}        // Начальные числа
	for i := 2; i < 10; i++ { // Генерируем числа
		next := fib[i-1] + fib[i-2] // Сумма двух прошлых чисел
		fib = append(fib, next)     // Добавляем следующее число в массив
	}
	fmt.Println(fib) // Вывод
}
