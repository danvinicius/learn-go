package main

import (
	"fmt"
	"time"
)

func escrever(txt string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", txt)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}

func main() {
	canal := escrever("Olá mundo") // encapsula a lógica da goroutine em uma função, não precisa lidar com isso na main
	for range 10 {
		fmt.Println(<-canal)
	}
}
