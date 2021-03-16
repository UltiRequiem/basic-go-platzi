package main

import "fmt"

func main() {
	myInterface := []interface{}{"Hola", 121, 4.90}
	fmt.Println(myInterface...)
}
