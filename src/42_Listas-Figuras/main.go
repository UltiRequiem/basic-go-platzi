package main

import (
	"fmt"
	"math"
	"reflect"
)

type figure2D interface {
	getArea() float64
}

type square struct {
	base float64
}

type rectangle struct {
	high  float64
	width float64
}

type trapezoid struct {
	baseA float64
	baseB float64
	high  float64
}

type circle struct {
	radio float64
}

func (s square) getArea() float64 {
	return s.base * s.base
}

func (r rectangle) getArea() float64 {
	return r.high * r.width
}

func (t trapezoid) getArea() float64 {
	return ((t.baseA + t.baseB) / 2) * t.high
}

func (c circle) getArea() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

func calculateArea(f figure2D) {
	fmt.Printf("Area of %s: %.2f\n", reflect.TypeOf(f).Name(), f.getArea())
}

func main() {
	mySquare := square{base: 5}
	myRectangle := rectangle{width: 4, high: 5}
	myTrapezoid := trapezoid{baseA: 18, baseB: 10, high: 5}
	myCircle := circle{radio: 4}
	calculateArea(mySquare)
	calculateArea(myRectangle)
	calculateArea(myTrapezoid)
	calculateArea(myCircle)
}
