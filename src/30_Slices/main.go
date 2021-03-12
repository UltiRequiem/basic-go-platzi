package main

import "fmt"

func main() {
	// Primer Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("Los valores del slice son:", slice, "\n", "Hay", len(slice), "elementos.", "\n", "La capacidad maxima del slice es", cap(slice), ".")
}
