package main

import "fmt"

func calcularMediaAluno(notas ...int) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	// DEFER = ADIAR -> Adia a execução para o final, imediatamente antes do retorno
	var soma int

	for _, nota := range notas {
		soma += nota
	}

	if soma/len(notas) >= 6 {
		return true
	}
	return false

}

func main() {
	res := calcularMediaAluno(7, 4, 9, 2, 5)
	fmt.Println(res)
}
