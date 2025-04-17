package main

import "fmt"

func main() {
	readers := map[string]map[string][]string{
		"Читатель1": {
			"Книги":         {"Книга1", "Книга2"},
			"Периодические": {"Журнал1"},
		},
		"Читатель2": {
			"Книги":         {"Книга3"},
			"Периодические": {},
		},
	}

	// Количество читателей
	fmt.Println("Количество читателей:", len(readers))

	// Общее количество изданий у каждого читателя
	for reader, publications := range readers {
		total := len(publications["Книги"]) + len(publications["Периодические"])
		fmt.Printf("%s: %d изданий\n", reader, total)
	}
}
