package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./hello.txt")
	if err != nil {
		fmt.Println("Hubo un error: " + err.Error())
	}

	index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		index++
		line := scanner.Text()
		fmt.Printf("%d: %s\n", index, line)
	}
}
