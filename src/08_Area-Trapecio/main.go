package main

import "fmt"

func main() {
	const baseUno int8 = 6
	const baseDos int8 = 15
	const alturaTrapecio int8 = 25

	const areaTrapecio int16 = ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Área del Trapecio es :", areaTrapecio)

}
