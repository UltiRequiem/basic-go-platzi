package main

import "fmt"

func main() {
	const one int = 1
	const two int = 2

	// Puerta Logica With And

	if one == 1 && two == 2 {
		fmt.Println("El primer valor es igual a uno y el segundo valor es igual a dos.")
	} else {
		fmt.Println("Puede que el primer valor no sea igual a uno y/o el segundo valor no es igua a dos.")
	}

}
