package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var a int
	fmt.Print("Введите ваше число - ")
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println("Ваше число четное")
	} else {
		fmt.Print("Число нечетное")
	}
}
