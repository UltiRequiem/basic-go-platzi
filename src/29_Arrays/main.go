package main

import "fmt"

func main() {
	// Primer Array
	var array [4]int
	// fmt.Println(array): Imprime el Zero Value

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	fmt.Println("Los valores del array son:", array, "\n", "Hay", len(array), "elementos.", "\n", "La capacidad maxima del array es", cap(array), ".")
}
