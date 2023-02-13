package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()
	channel := make(chan string)
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0
	for {
		if i > 2 {
			break
		}
		for _, server := range servers {
			go checkServer(server, channel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
		i++
	}
	// for i := 0; i < len(servers); i++ {
	// 	fmt.Println(<-channel)
	// }
	elapsedTime := time.Since(startTime)
	fmt.Printf("Tiempo ejecución %s\n", elapsedTime)
}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		channel <- server + " no esta disponible!"
	} else {
		channel <- server + " funcionando normalmente!"
	}
}
