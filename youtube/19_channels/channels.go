package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)

	go func(ch chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			ch <- name
		}
	}(channel)

	for {
		msg := <-channel
		fmt.Println("Mensaje: " + msg)
	}

}
