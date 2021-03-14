package main

import "fmt"

func main() {
	slice := []string{"Hola", "Mundo."}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}
}
