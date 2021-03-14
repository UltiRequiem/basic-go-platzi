package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
