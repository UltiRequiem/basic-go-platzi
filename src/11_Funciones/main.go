package main

import "fmt"

func holaMundo(number string) {
	fmt.Println("Hola Mundo", number)
}

func returnValue(a int) int {
	return a * 2
}

func main() {
	holaMundo("1")

	var value int = returnValue(75)
	fmt.Println("Value:", value)
}
