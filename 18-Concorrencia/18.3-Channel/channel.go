package main

import (
	"fmt"
	"time"
)

func escrever(txt string, canal chan string) {
	for range 5 {
		canal <- txt
		time.Sleep(time.Second)
	}

	close(canal) // fecha o canal pra ele não ficar esperando depois da última execução
}

func main() {
	canal := make(chan string)
	go escrever("Olá mundo", canal)

	fmt.Println("Depois de a função escrever começar a ser executada")

	// msg := <-canal // o programa espera o canal receber um dado para continuar a execução

	// for {
	// 	msg, open := <-canal // se não parar de escutar o canal, as goroutines vão terminar de executar e ele dá deadlock pq iria ficar esperando eternamente

	// 	if !open {
	// 		break
	// 	}

	// 	fmt.Println(msg)
	// }

	for msg := range canal { // sintaxe mais fácil e ele só roda enquanto o canal estiver aberto
		fmt.Println(msg)
	}
}
