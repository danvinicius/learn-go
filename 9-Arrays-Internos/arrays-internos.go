package main

import "fmt"

func main() {
	slice1 := make([]float32, 10, 11)
	// o make cria um aray interno de tamanho (capacidade max) 11, e retorna um slice de tamanho 10
	fmt.Println(len(slice1)) // tamanho
	fmt.Println(cap(slice1)) // capacidade

	slice1 = append(slice1, 42)
	slice1 = append(slice1, 73)

	// quando o slice estoura o tamanho do array, o go cria um novo array maior e o slice passa a apontar para ele

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
