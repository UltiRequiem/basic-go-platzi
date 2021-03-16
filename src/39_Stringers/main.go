package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco y es un sistema %s.", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 16, brand: "linux", disk: 100}

	fmt.Println(myPC)
}
