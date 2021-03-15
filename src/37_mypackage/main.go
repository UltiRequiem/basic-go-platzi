package main

import (
	"fmt"

	"github.com/UltiRequiem/Curso-Basico-de-Programacion-en-Go-de-Platzi/go/src/Curso-Basico-de-Programacion-en-Go-de-Platzi/src/37_mypackage/car"
)

func main() {
	var myCar car.Car
	myCar.Brand = "Ferrari"
	myCar.Year = 2020

	fmt.Println(myCar)

	car.PrintMessage("Funcion Publica")

	car.printMessage("Funcion Privada")

	//	var myOtherCar car.carPrivate
	//	fmt.Println(myOtherCar)
}
