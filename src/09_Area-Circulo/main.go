package main

import "fmt"

func main() {
	// Defino el valor de pi
	const pi = 3.14
	var radioCirculo float64 = 50

	areaCirculo := pi * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

}
