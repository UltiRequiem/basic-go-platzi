package main

import "fmt"

func main() {
	// Defino el valor de pi
	const pi float32 = 3.14
	const radioCirculo float32 = 50

	const areaCirculo float32 = pi * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

}
