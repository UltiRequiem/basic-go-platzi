package main

import "fmt"

func main() {
	const baseUno uint16 = 6
	const baseDos uint16 = 15
	const alturaTrapecio uint16 = 25

	const areaTrapecio uint16 = ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Área del Trapecio es :", areaTrapecio)

}
