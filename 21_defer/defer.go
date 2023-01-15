package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./hello.txt")

	// Defer permite ejecutar rutinas aunque cambie el flujo normal del código
	// similar al funcionamiento de "finally"
	defer func() {
		file.Close()
		fmt.Println("Archivo cerrado con éxito!")
	}()

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
