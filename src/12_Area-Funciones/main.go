package main

import "fmt"

func AreaRectangulo(baseRectangulo, alturaRectangulo int) {
	var areaRectangulo int = baseRectangulo * alturaRectangulo
	fmt.Println(areaRectangulo)
}

func AreaTrapeze(baseMenor, baseMayor, altura int) {
	var areaTrapeze int = (baseMenor + baseMayor) * altura / 2

	fmt.Println(areaTrapeze)
}

func areaCircle(radio float64) {
	const pi float64 = 3.14
	var areaCircle float64 = pi * radio * radio

	fmt.Println(areaCircle)
}

func main() {

	AreaRectangulo(10, 20)

	AreaTrapeze(15, 7, 25)

	areaCircle(50)

}
