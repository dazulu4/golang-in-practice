package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Hola mundo")

	// edad := 22
	// fmt.Printf("Mi edad es %d\n", edad)
	// fmt.Printf("Mi edad es %v\n", edad)

	// nombre := "Uriel"
	// fmt.Printf("Mi nombre es %s\n", nombre)
	// fmt.Printf("Mi nombre es %v\n", nombre)

	// bandera := true
	// fmt.Printf("Mi bandera es %t\n", bandera)

	// precio := 14.33535635346
	// fmt.Printf("Mi precio es %f\n", precio)
	// fmt.Printf("Mi precio es %.2f\n", precio)

	// var edad_in int
	// fmt.Print("Ingresa la edad: ")
	// fmt.Scanf("%d\n", &edad_in)
	// fmt.Printf("La edad es: %d\n", edad_in)

	// var nombre_in string
	// fmt.Print("Ingresa el nombre: ")
	// fmt.Scanf("%s\n", &nombre_in)
	// fmt.Printf("El nombre es %s\n", nombre_in)

	reader := bufio.NewReader((os.Stdin))
	fmt.Println("Ingresa tu nombre: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + name)
	}
}
