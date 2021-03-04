package main

import "fmt"

func main() {
	//Delcarando variables enteras:
	var base int16 = 12
	var altura int16 = 14

	// Posibles formas de Declarar esta variable
	// var area = altura * base
	// area := altura * base

	// Forma correcta de hacerlo:
	var area int16 = altura * base

	fmt.Println("La base vale:", base, "\nLa altura vale:", altura, "\nY el area es igual a:", area)

}
