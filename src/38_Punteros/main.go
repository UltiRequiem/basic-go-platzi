package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

func (myPC *pc) duplicateDisk() {
	myPC.disk = myPC.disk * 2
}

func main() {
	var a int = 50
	b := &a

	fmt.Println(*b, b)

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "linux"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRam()
	myPC.duplicateDisk()

	fmt.Println(myPC)

}
