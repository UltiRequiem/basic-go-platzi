package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Adios")
	}()

	time.Sleep(time.Second * 1)
}
