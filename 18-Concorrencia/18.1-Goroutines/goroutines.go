package main

import (
	"fmt"
	"time"
)

func escrever(txt string) {
	for {
		time.Sleep(time.Second)
		fmt.Println(txt)
	}
}

func main() {
	go escrever("Olá mundo")
	escrever("Programando em Go")
}
