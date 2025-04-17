package main

import "fmt"

func main() {
	var B int
	A := &B
	*A = 35
	fmt.Println(B)
}
