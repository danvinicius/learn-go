package main

import (
	"fmt"
	"sync"
	"time"
)

type BancoDeDados struct {
	conectado bool
}

var (
	bancoDeDados *BancoDeDados
	once         sync.Once
)

func ConectarAoBanco() *BancoDeDados {
	once.Do(func() { // tudo que tiver aqui dentro só será executado uma única vez (precisa compartilhar a mesma variável once), como se fosse um singleton
		fmt.Println("Conectando ao banco de dados...")
		time.Sleep(time.Second * 2)
		bancoDeDados = &BancoDeDados{conectado: true}
		fmt.Println("Conectado!")
	}) // além disso, também vai bloquear para que uma goroutine execute por vez

	return bancoDeDados
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := range 5 {
		go func() {
			defer wg.Done()
			fmt.Printf("Goroutine %d tentando conectar ao banco de dados\n", i)
			bancoDeDados = ConectarAoBanco()
			fmt.Printf("Goroutine %d finalizando. Conectado = %v\n", i, bancoDeDados.conectado)
		}()
	}

	wg.Wait()
}
