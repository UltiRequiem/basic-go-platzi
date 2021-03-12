package main

import "fmt"

func main() {
	var value int = -98
	switch {
	case value > 100:
		fmt.Println("Es un número mayor a 100.")
	case value < 0:
		fmt.Println("Es un número menor a 0.")
	default:
		fmt.Println("No cumple ninguna condición.")
	}
}
