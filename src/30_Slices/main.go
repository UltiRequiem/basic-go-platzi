package main

import "fmt"

func main() {
	// Primer Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	var message string = fmt.Sprint("Los valores del slice son:", slice, "\n", "Hay ", len(slice), " elementos.", "\n", "La capacidad maxima del slice es ", cap(slice), ".")
	fmt.Println(message)

	fmt.Println(slice[:3])
	fmt.Println(slice[4:], "\n")

	slice = append(slice, 7)
	fmt.Println("Ahora Slice contiene", slice)

}
