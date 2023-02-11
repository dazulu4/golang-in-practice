package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer el map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Encontrar un valor
	value := m["Jose"]
	fmt.Println(value)

	value2, ok := m["Joseph"]
	fmt.Println(value2, ok)
}
