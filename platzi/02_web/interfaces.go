package main

import "fmt"

type animal interface {
	move() string
}

type dog struct {
}

type fish struct {
}

type bird struct {
}

func (dog) move() string {
	return "Soy un perro y camino"
}

func (fish) move() string {
	return "Soy un pez y nado"
}

func (bird) move() string {
	return "Soy un pajaro y vuelo"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	d := dog{}
	moveAnimal(d)

	f := fish{}
	moveAnimal(f)

	b := bird{}
	moveAnimal(b)
}
