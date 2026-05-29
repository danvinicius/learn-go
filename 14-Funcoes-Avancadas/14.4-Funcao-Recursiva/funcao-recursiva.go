package main

import "fmt"

func fibonnacci(pos uint) uint {
	if pos <= 2 {
		return 1
	}
	return fibonnacci(pos-1) + fibonnacci(pos-2)
}

func main() {
	fib := fibonnacci(5)
	fmt.Println(fib)
}
