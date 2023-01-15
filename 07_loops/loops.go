package main

import "fmt"

func main() {
	for i := 1; i < 10; i = i + 2 {
		fmt.Println(i)
	}

	//While
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//Infinite While
	j := 0
	for {
		j++
		if j == 2 {
			continue
		}
		fmt.Println(j)
		if j >= 10 {
			break
		}
	}
}
