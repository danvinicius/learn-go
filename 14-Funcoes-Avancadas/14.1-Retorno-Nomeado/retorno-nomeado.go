package main

import "fmt"

func calcula(x int, y int) (soma int, sub int) {
	soma = x + y
	sub = x - y
	return
}

func main() {
	var soma, sub = calcula(8, 7)
	fmt.Println(soma, sub)
}
