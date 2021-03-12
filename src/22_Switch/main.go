package main

import "fmt"

func main() {
	var modulo int = 6 % 2
	switch modulo {
	case 0:
		fmt.Println("Es un número par.")
	default:
		fmt.Println("Es un número impar.")
	}
}
