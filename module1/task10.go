package main

import "fmt"

func main() {
	num := []int{1, 2, 5, 1, 3, 5, 6} // Создаем массив
	uniqle := []int{}                 // Для вложения уникальных чисел
	// Проверяем число на уникальность
	for _, i := range num {
		trUniqle := true
		for _, j := range uniqle {
			if i == j {
				trUniqle = false
				break
			}
		}
		if trUniqle {
			uniqle = append(uniqle, i)
		}
	}
	fmt.Println(uniqle)
}
