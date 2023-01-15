package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file_data, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		fmt.Println("Hubo un error: " + err.Error())
	}

	fmt.Println(string(file_data))
}
