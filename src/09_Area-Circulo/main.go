package main

import "fmt"

func main() {
	// Defino el valor de pi
	const pi float64 = 3.14
	const radioCirculo float64 = 50

	const areaCirculo float64 = pi * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

}
