package main

import "strconv"
import "fmt"

func main() {
	str := "104"
	num := 35
	convStr, _ := strconv.Atoi(str)
	convInt := strconv.Itoa(num)
	fmt.Println(convStr)
	fmt.Println(convInt)
}
