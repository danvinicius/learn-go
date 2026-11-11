package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(i int, wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Println("Iniciando o worker", i)
	time.Sleep(time.Second)
	if i == 2 {
		return errors.New("erro inesperado")
	}
	fmt.Println("Finalizando o worker", i)
	return nil
}

func workerWithErrorGroup(i int) error {
	fmt.Println("Iniciando o worker", i)
	time.Sleep(time.Second)
	if i == 2 {
		return errors.New("erro inesperado")
	}
	fmt.Println("Finalizando o worker", i)
	return nil
}

func main() {

	fmt.Println("=============== WAIT GROUP (tratamento de erros convencional/manual) ===============")
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		wg.Go(func() {
			err := worker(i, &wg)
			if err != nil {
				fmt.Println("Houve um erro no item ", i)
			}
		})
	}

	wg.Wait()

	fmt.Println("=============== ERROR GROUP (tratamento de erros nativo) ===============")

	var eg errgroup.Group

	for i := range 3 {
		eg.Go(func() error {
			return workerWithErrorGroup(i)
		})
	}

	err := eg.Wait() // ele armazena o primeiro erro
	if err != nil {
		fmt.Println("Houve um erro:", err)
	}
}
