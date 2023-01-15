package main

import "fmt"

type Human struct {
	name string
}

func (this Human) talk() string {
	return "Bla bla bla"
}

type Dummy struct {
	name string
}

type Tutor struct {
	// Campo anonimo, me permite hacer "herencia"
	Human
	Dummy
}

func (this Tutor) talk() string {
	return "Bienvenidos a código facilito"
}

func main() {
	tutor := Tutor{Human{"Uriel"}, Dummy{"Alex"}}

	// fmt.Println(tutor.name) Esto marca error
	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.Dummy.name)

	// Utiliza la implementación
	fmt.Println(tutor.talk())
}
