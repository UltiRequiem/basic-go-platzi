package main

import (
	"Curso-Basico-de-Programacion-en-Go-de-Platzi/src/37_mypackage/car"
	"fmt"
)

func main() {
	var myCar car.Car
	myCar.Brand = "Ferrari"
	myCar.Year = 2020

	fmt.Println(myCar)

	car.PrintMessage("Funcion Publica")

	//car.printMessage("Funcion Privada")
}
