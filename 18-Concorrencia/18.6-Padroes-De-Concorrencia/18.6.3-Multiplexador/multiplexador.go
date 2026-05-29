package main

import (
	"fmt"
	"time"
)

func multiplexar(canal1, canal2 <-chan string) <-chan string { // junta dois canais em um
	canalFinal := make(chan string)
	go func() {
		for {
			select {
			case msg := <-canal1:
				canalFinal <- msg
			case msg := <-canal2:
				canalFinal <- msg
			}
		}
	}()
	return canalFinal
}

func main() {
	canal := multiplexar(escrever("Olá mundo"), escrever("Programando em Go"))
	for range 10 {
		fmt.Println(<-canal)
	}
}

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
