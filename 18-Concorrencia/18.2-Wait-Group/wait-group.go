package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(txt string) {
	for range 5 {
		time.Sleep(time.Second)
		fmt.Println(txt)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // add 2 goroutines numa "fila"

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // espera as 2 goroutines serem executadas e continua o programa

	fmt.Println("continuando o programa ...")

	// se eu só chamasse 2 goroutines, ele iria pular as duas funções e terminar o programa sem executar nada
}
