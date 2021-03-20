package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println("En el canal hay", len(c), ". Y su maxima capacidad es ", cap(c), ".")
}
