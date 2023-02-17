package main

import (
	"fmt"
	"sync"
	"time"
)

func doSometing(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Ended %d\n", i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSometing(i, &wg)
	}
	wg.Wait()
}
