package main

import "fmt"

func main() {
	// Valores:
	var x uint16 = 100
	var y uint16 = 50

	//Suma:
	var sumaResult uint16 = x + y
	//Resta:
	restaResult := x - y

	fmt.Println("La suma de los 2 valores es ", sumaResult, ".Mientras que su resta equivale a ", restaResult, ".")

	//Multiplicacion:
	var multiResult uint16 = x * y
	// División:
	var diviResult uint16 = x / y

	fmt.Println("La multiplicación de los dos valores es", multiResult, " y su divisón es igual a", diviResult, ".")

	// Modulo:
	var modulResult uint16 = x % y

	fmt.Println("El modulo de los dos valores es: ", modulResult)

	// Incremental:
	x++

	fmt.Println("El incremental de x es: ", x)
	// Decremental:
	x--
	fmt.Println("El decremental de x es: ", x)

}
