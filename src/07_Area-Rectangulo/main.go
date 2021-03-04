package main

import "fmt"

func main() {
	// Información de mi rectángulo:
	const ladoRectangulo uint16 = 100
	const altoRectangulo uint16 = 20

	const areaRectangulo uint16 = ladoRectangulo * altoRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

}
