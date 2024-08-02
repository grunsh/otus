package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func retReversString(s string) string {
	return reverse.String(s)
}

func main() {

	exampleString := "Hello, OTUS!"

	fmt.Println(retReversString(exampleString))

}
