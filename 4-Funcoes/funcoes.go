package main

import "fmt"

func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1 int8, n2 int8) (int8, int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	mult := n1 * n2

	return soma, sub, mult
}

func main() {
	soma1 := soma(7, 5)
	soma2 := soma(3, 3)
	soma, sub, _ := calculos(30, 15)
	fmt.Println(soma, sub)

	var mult = func(n1 int8, n2 int8) int8 {
		return n1 * n2
	}

	x := mult(soma1, soma2)
	fmt.Println(x)
}
