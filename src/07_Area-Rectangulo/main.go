package main

import "fmt"

func main() {
	// Información de mi rectángulo:
	const ladoRectangulo int16 = 100
	const altoRectangulo int16 = 20

	const areaRectangulo int16 = ladoRectangulo * altoRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

}
