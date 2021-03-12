package main

import "fmt"

func main() {
	if isPair(7) {
		fmt.Println("Number is pair.")
	} else {
		fmt.Println("Number is odd.")
	}
}

func isPair(num int) bool {
	return num%2 == 0
}
