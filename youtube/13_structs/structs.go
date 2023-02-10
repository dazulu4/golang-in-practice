package main

import "fmt"

// Con el keyword "type" indica que se definirá un nuevo tipo de dato que es una estructura
type User struct {
	age            int
	name, lastname string
}

func main() {
	var uriel User
	fmt.Println(uriel)

	fmt.Println(User{name: "Uriel", lastname: "Hernandez"})

	fmt.Println(User{42, "Alexander", "Zuluaga"}.name)

	uriel2 := new(User)
	fmt.Println(uriel2.name)

	// Las estructuras son mutables, podemos modificar los valores
	uriel2.name = "Uriel"
	fmt.Println(uriel2.name)
}
