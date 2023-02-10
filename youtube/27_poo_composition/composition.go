package main

import "fmt"

type liquid struct {
	color string
}

type brand struct {
	country string
}

func (l *liquid) isColor() {
	l.color = "white"
	fmt.Println(l.color)
}

// func (l liquid) isColorNoRef() {
// 	l.color = "white"
// 	fmt.Println(l.color)
// }

type beer struct {
	name    string
	alcohol float32
	liquid
	brand
}

func (b *beer) isBeer() {
	fmt.Println("Mi cerveza es", b.name, b.color, b.alcohol, b.country)
}

func main() {
	l := liquid{color: "yellow"}
	// l.isColor()
	// fmt.Println(l.color)

	// l2 := liquid{color: "yellow"}
	// l2.isColorNoRef()
	// fmt.Println(l2.color)

	br := brand{country: "Alemania"}

	b := beer{name: "Aguila", alcohol: 4.8, liquid: l, brand: br}
	b.isBeer()
}
