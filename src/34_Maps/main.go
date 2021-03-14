package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	//fmt.Println(m)

	// Recorrer map

	for index, value := range m {
		fmt.Println(index, value)
	}

}
