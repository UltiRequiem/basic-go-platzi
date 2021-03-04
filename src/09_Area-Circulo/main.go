package main

import "fmt"

func main() {
	// Defino el valor de pi
	const pi ufloat32 = 3.14
	const radioCirculo ufloat32 = 50

	const areaCirculo ufloat32 = pi * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

}
