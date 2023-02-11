package main

import (
	"21_access_modifiers/mypackage"
	"fmt"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	// var myOtherCar mypackage.carPrivate
	// fmt.Println(myOtherCar)

	mypackage.PrintMessage("Hola Platzi")

	// mypackage.printMessage1("Hola Platzi Private")
}
