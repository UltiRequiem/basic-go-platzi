package main

import "fmt"

func calcularAreaRectangulo(baseRectangulo, alturaRectangulo int) {
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println(areaRectangulo)
}

func calcularAreaTrapecio(baseMenor, baseMayor, altura int) {
	areaTrapecio := (baseMenor + baseMayor) * altura / 2
	fmt.Println(areaTrapecio)
}

func calcularAreaCirculo(radio float64) {
	const pi float64 = 3.14
	areaCirculo := pi * radio * radio

	fmt.Println(areaCirculo)
}

func main() {

	calcularAreaRectangulo(5, 10)

	calcularAreaTrapecio(5, 10, 3)

	calcularAreaCirculo(10)

}
