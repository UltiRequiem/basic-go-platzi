package main

import "fmt"

func main() {
	const one int = 14
	const two int = 27

	// Puerta Logica With Or

	if one == 1 || two == 2 {
		fmt.Println("El primer valor es igual a uno o el segundo valor es igual a dos.")
	} else {
		fmt.Println("El primer valor no es igual a uno y el segundo valor no es igua a dos.")
	}

}
