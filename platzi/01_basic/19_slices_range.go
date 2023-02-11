package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReserve string
	for i := len(text) - 1; i >= 0; i-- {
		textReserve += string(text[i])
	}
	if strings.ToLower(text) == strings.ToLower(textReserve) {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	// slice := []string{"hola", "que", "hace"}

	// for i, valor := range slice {
	// 	fmt.Println(i, valor)
	// }

	// for _, valor := range slice {
	// 	fmt.Println(valor)
	// }

	// for j := range slice {
	// 	fmt.Println(j)
	// }

	isPalindromo("ama")
	isPalindromo("Ama")
	isPalindromo("casa")

}
