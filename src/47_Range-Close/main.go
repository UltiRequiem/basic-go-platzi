package main

import (
	"fmt"
)

func mensaje(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	close(c) // Cierra el Canal

	//c <- "Mensaje 3" Nos daria eror de runtime

	for mensaje := range c {
		fmt.Println(mensaje)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)
	go mensaje("mensaje1", email1)
	go mensaje("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1:", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2: ", m2)
		}
	}
}
