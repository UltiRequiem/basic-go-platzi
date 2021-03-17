package main

import (
	"fmt"
	"time"
)

func main() {
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}
