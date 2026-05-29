package main

import "fmt"

func inverteSinal(num int) int {
	return num * -1
}

func inverteComPonteiro(num *int) int {
	invertido := *num * -1
	*num = invertido
	return invertido
}

func main() {
	num := 20
	invertido := inverteSinal(num)

	fmt.Println(invertido)
	fmt.Println(num) // se mantém inalterado pois é uma cópia

	invertidoComPonteiro := inverteComPonteiro(&num)
	fmt.Println(invertidoComPonteiro)
	fmt.Println(num) // alterei o valor real que estava no endereço de memória
}
