package main

import (
	"fmt"
	"strings"
	"time"
)

func slow_method(name string) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(letter)
	}
}

func main() {
	// slow_method("Uriel")
	// fmt.Println("Que aburrido esperar!")

	// Hagamoslo con Go Routine
	go slow_method("Uriel")
	fmt.Println("Que aburrido esperar!")

	// var wait string
	// fmt.Scanln(&wait)

	go func() {
		var wait string
		fmt.Scanln(&wait)
	}()

}
