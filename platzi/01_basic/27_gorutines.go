package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string) {
	fmt.Println(text)
}

func saySync(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	// Mala práctica
	// say("Hello")
	// go say("World")
	// time.Sleep(time.Second * 1)

	var wg sync.WaitGroup
	say("Hello")
	wg.Add(1)
	go saySync("World", &wg)
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")
	time.Sleep(time.Second * 1)
}
