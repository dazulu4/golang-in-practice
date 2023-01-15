package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	flag := read_file()
	fmt.Printf("%v\n", flag)
	fmt.Println("Finaliza ejecución programa!")
}

func read_file() bool {
	file, err := os.Open("./hellooo.txt")
	defer func() {
		file.Close()
		fmt.Println("Archivo cerrado con éxito!")

		rec := recover()
		fmt.Println(rec)
	}()

	// Panic permite imprimir el la traza de errores generado
	// similar a como funciona "traceback" o "stack_trace"
	// retorna todas las funciones anidadas con un código de error hasta el main
	// relanza la excepcion algo similar a "throw exception"
	if err != nil {
		panic(err)
	}

	index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		index++
		line := scanner.Text()
		fmt.Printf("%d: %s\n", index, line)
	}

	return true
}
