package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

// Реализуем в виде своей функции, вдруг потом надо будет дополнить фичами.
func retReversString(s string) string {
	return reverse.String(s)
}

func main() {
	// Строка для перевёртыша
	exampleString := "Hello, OTUS!"
	// Выводим в стандартный вывод вывернутую строку
	fmt.Println(retReversString(exampleString))
}
