package main

import "fmt"

func main() {
	var str string
	reversedStr := ""
	fmt.Print("Введите вашу строку - ")
	fmt.Scan(&str)
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}

	fmt.Println(reversedStr)
}
