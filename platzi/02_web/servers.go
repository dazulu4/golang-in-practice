package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	for _, server := range servers {
		checkServer(server)
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("Tiempo ejecución %s\n", elapsedTime)
}

func checkServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "no esta disponible!")
	} else {
		fmt.Println(server, "funcionando normalmente")
	}
}
