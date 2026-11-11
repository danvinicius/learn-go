package main

import (
	"fmt"
	"time"
)

func send(ch chan <- int) {
	fmt.Println("enviando valor ao channel")
	ch <- 100 // se não for um channel com buffer, é uma operação bloqueante
	fmt.Println("valor enviado ao channel")
}

func receive(ch <- chan int) {
	time.Sleep(time.Second)
	valor := <-ch
	fmt.Println("valor recebido pelo channel: ", valor)
}

// quando usar channels -> comunicação entre goroutines
func main() {
	ch := make(chan int)

	for range 3 {
		go send(ch)
	}

	time.Sleep(time.Second)

	for range 3 {
		go receive(ch)
	}

	time.Sleep(time.Second * 3)

	// <- ch deadlock!

	chComBuffer := make(chan int, 3) // com buffer de 3, ele não irá bloquear até 3 mensagens, se passar disso ele se comporta como um channel normal

	for range 3 {
		go send(chComBuffer)
	}

	time.Sleep(time.Second)

	for range 3 {
		go receive(chComBuffer)
	}

	time.Sleep(time.Second * 3)
}
