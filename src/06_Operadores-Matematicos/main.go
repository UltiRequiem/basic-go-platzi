package main

import "fmt"

func main() {
	// Valores:
	x := 1000
	y := 500

	//Suma:
	sumaResult := x + y
	//Resta:
	restaResult := x - y

	fmt.Println("La suma de los 2 valores es ", sumaResult, ".Mientras que su resta equivale a ", restaResult, ".")

	//Multiplicacion:
	multiResult := x * y
	// División:
	diviResult := x / y

	fmt.Println("La multiplicación de los dos valores es", multiResult, " y su divisón es igual a", diviResult, ".")

	// Modulo:
	modulResult := x % y

	fmt.Println("El modulo de los dos valores es: ", modulResult)

	// Incremental:
	x++
	fmt.Println("El incremental de x es: ", x)
	// Decremental:
	x--
	x--
	fmt.Println("El decremental de x es: ", x)

}
