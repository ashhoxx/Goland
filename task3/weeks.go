package main

import "fmt"

func main() {
	days := []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье"}
	var work []string

	for _, day := range days {
		if day != "Суббота" && day != "Воскресенье" {
			work = append(work, day)
		}
	}

	weekdays := []string{"Суббота", "Воскрсенье"}

	fmt.Println("Будние дни - ", work)
	fmt.Println("Выходные дни - ", weekdays)
}
