package main

import "fmt"

func main() {
	const baseUno uint8 = 6
	const baseDos uint8 = 15
	const alturaTrapecio uint8 = 25

	const areaTrapecio uint16 = ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Área del Trapecio es :", areaTrapecio)

}
