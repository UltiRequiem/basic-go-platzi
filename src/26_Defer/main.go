package main

import "fmt"

func main() {
	defer fmt.Println("Ejecutado al final pese a estar al comienzo gracias a 'Defer'.")
	fmt.Println("Hola Mundo")
}
