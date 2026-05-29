package main

import (
	"fmt"
	"time"
)

func fibonnacci(pos uint) uint {
	if pos <= 2 {
		return 1
	}
	return fibonnacci(pos-1) + fibonnacci(pos-2)
}

func worker(tarefas <-chan uint, resultados chan<- uint) {
	for numero := range tarefas {
		resultados <- fibonnacci(numero)
	}
}

func main() {
	var tam uint = 45

	tarefas := make(chan uint, tam)
	resultados := make(chan uint, tam)

	inicio := time.Now()

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := uint(0); i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := uint(0); i < 45; i++ {
		num := <-resultados
		fmt.Println(num)
	}

	duracao := time.Since(inicio)

	fmt.Printf("Tempo de execução: %v\n", duracao)
}

// 1 worker: 5.272403397s
// 2 workers: 3.440695621s
// 3 workers: 2.943645195s
// 4 workers: 2.687574157s
// 5 workers: 2.495924105s
