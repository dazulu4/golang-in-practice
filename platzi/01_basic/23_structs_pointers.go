package main

import "fmt"

type pc struct {
	ram   int
	dick  int
	brand string
}

func (pcObj pc) ping() {
	fmt.Println(pcObj.brand, "Pong")
}

func (pcObj *pc) duplicateRam() {
	pcObj.ram = pcObj.ram * 2
}

func (pcObj pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB disco y es una %s", pcObj.ram, pcObj.dick, pcObj.brand)
}

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, dick: 250, brand: "msi"}
	myPC.ping()

	fmt.Println(myPC)

	myPC.duplicateRam()
	fmt.Println(myPC)

	myPC.duplicateRam()
	fmt.Println(myPC)
}
