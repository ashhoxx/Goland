package main

import "fmt"

func main() {
	weekend := []string{"Суббота", "Воскресенье"}
	weekday := []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница"}

	combo := append(weekday, weekend...)

	fmt.Println(combo)

}
