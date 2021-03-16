package main

import "fmt"

// Mi Interfaz
type figuras2D interface {
	area() float64
}

// Mis Estructuras
type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

// Mis funciones
func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

// Main()
func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 4, altura: 8}

	calcular(myCuadrado)   // 2*2
	calcular(myRectangulo) // 4*8
}
