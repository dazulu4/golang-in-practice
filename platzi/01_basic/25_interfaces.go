package main

import "fmt"

type figura2d interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figura2d) {
	fmt.Println("Area:", f.area())
}

func main() {
	miCuadrado := cuadrado{base: 2}
	miRectangulo := rectangulo{base: 2, altura: 4}
	calcular(miCuadrado)
	calcular(miRectangulo)

	// Lista de interfaces
	miInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(miInterface...)
}
