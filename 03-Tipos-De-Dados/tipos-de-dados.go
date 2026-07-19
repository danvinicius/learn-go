package main

import (
	"errors"
	"fmt"
)

func main() {
	// NÚMEROS
	// int e float, que se dividem em várias subcategorias

	// int8, int16, int32, int64 // qtd de bits
	// uint8, uint16, uint32, uint64 // mesma coisa mas sem sinal
	// exemplo: var x int8 = 128 da erro
	// int utiliza arq do computador como base, pc 64 bits = int64, pc 32 bits = int32

	var num int16 = 10000
	fmt.Println(num)

	// alias
	var num1 rune = 12345
	fmt.Println(num1)

	var num2 byte = 123
	fmt.Println(num2)

	var numeroReal1 float32 = 1234.56
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234.56
	fmt.Println(numeroReal2)

	// float (não pode declarar explicitamente) tb utiliza arq do computador como base

	// STRING

	string1 := "string"
	fmt.Println(string1)

	char := 'B'
	fmt.Println(char) // converte pra numero da tabela ASCII, não existe char na prática no go

	// BOOLEANO
	var bol1 bool = true
	fmt.Println(bol1)

	var erro1 error
	fmt.Println(erro1)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

	// no Go, todo tipo de dado tem seu valor zero / valor inicial, quando eu inicializo uma variável mas não declaro o valor
	// numero -> 0
	// string -> ""
	// bool -> false
	// erro -> nil

}
