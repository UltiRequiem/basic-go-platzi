package main

import "fmt"

type car struct {
	brand   string
	year    int
	seating int
	color   string
	owner   string
}

func main() {
	myCar := car{brand: "Toyota", year: 2018, seating: 10, color: "Rojo", owner: "Eliaz Bobadilla"}
	fmt.Println("Los Datos de mi auto son:", myCar)

	var myOtherCar car
	myOtherCar.brand = "Ford"
	fmt.Println(myOtherCar)

}
