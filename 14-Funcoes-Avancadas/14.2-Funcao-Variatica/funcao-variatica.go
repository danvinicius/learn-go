package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}

func main() {
	soma := soma(1, 2, 3, 4)
	fmt.Println(soma)
}
