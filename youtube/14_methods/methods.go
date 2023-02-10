package main

import "fmt"

type User struct {
	age            int
	name, lastname string
}

func (this User) getCompleteName() string {
	return this.name + " " + this.lastname
}

// Pasa "this" como una copia del objeto
func (this User) setName(name string) {
	this.name = name
}

func (this *User) setRealName(name string) {
	this.name = name
}

func main() {
	uriel := new(User)
	uriel.name = "Uriel"
	uriel.lastname = "Hernandez"
	fmt.Println(uriel.getCompleteName())

	// Estamos pasando objeto uriel como una copia
	uriel.setName("Marcos")
	fmt.Println(uriel.getCompleteName())

	// Estamos pasando el puntero del objeto uriel
	uriel.setRealName("Marcos")
	fmt.Println(uriel.getCompleteName())
}
