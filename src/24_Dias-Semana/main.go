package main

import "fmt"

func main() {
	var dia int8 = 4

	switch dia {
	case 1:
		fmt.Println("Hoy es Lunes")
	case 2:
		fmt.Println("Hoy es Martes")
	case 3:
		fmt.Println("Hoy es Miercoles")
	case 4:
		fmt.Println("Hoy es Jueves")
	case 5:
		fmt.Println("Hoy es Viernes")
	case 6:
		fmt.Println("Hoy es Sabado")
	case 7:
		fmt.Println("Hoy es Domingo")
	default:
		fmt.Println("Ese no es un día valido de la semana!")
	}
}
