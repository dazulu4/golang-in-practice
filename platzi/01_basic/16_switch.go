package main

import "fmt"

func main() {
	// Normal
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// En una línea
	switch modulo1 := 4 % 2; modulo1 {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condición
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor que 100")
	case value < 0:
		fmt.Println("Es menor que 0")
	default:
		fmt.Println("No condición")
	}
}
