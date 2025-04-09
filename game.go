package main

import (
	"fmt"
	"math/rand" // Для генерации случайных чисел
	"time"      // Для работы со временем
)

func main() {
	// Создание нового генератора случайных чисел
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	secretNumber := r.Intn(100) + 1 // Загаданное число от 1 до 100
	var guess int                   // Для хранение введеных чисел
	attempts := 0                   // Счетчик попыток

	fmt.Println("Добро пожаловать в игру 'Угадай число'!")
	fmt.Println("Я загадал число от 1 до 100. Попробуйте угадать его!")

	for { // Бесконечный цикл с числами введеных пользователем
		fmt.Print("Введите ваше предположение: ")
		_, err := fmt.Scan(&guess)

		if err != nil { // Если число выходит за диапазон
			fmt.Println("Пожалуйста, введите валидное число.")
			continue
		}

		attempts++ // Прибавляем попытки

		if guess < secretNumber {
			fmt.Println("Слишком низко! Попробуйте еще раз.")
		} else if guess > secretNumber {
			fmt.Println("Слишком высоко! Попробуйте еще раз.")
		} else {
			fmt.Printf("Поздравляю! Вы угадали число %d за %d попыток.\n", secretNumber, attempts)
			break
		}
	}
}
