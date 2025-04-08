package main

import "fmt"

func isPalindrome(s string) bool {
	reversedStr := ""
	for _, char := range s {
		reversedStr = string(char) + reversedStr
	}
	return s == reversedStr
}
func main() {
	var word string
	fmt.Print("Введите ваше слово - ")
	fmt.Scan(&word)
	if isPalindrome(word) {
		fmt.Print("Ваша строка - палидром")
	} else {
		fmt.Print("Ваша строка - не палидром")
	}
}
