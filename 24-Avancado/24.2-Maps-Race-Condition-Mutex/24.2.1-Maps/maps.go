package main

import "fmt"

func main() {
	var meuMap = map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}

	/*

		+----------+----------------------+
		| Bucket 0 | [{"a": 1}, {"e": 5}] |
		+----------+----------------------+
		| Bucket 1 | [{"b": 2}]     			|
		+----------+----------------------+
		| Bucket 2 | [{"c": 3}, {"f": 6}] |
		+----------+----------------------+
		| Bucket 3 | [{"d": 4}]    		  	|
		+----------+----------------------+

		O go passa os campos do map para uma função de hash e guarda esses valores nos buckets da forma mais eficiente
		Os buckets podem mudar a cada execução
	*/

	// o range vai iterar através dos buckets, então não é garantida a ordem
	for k, v := range meuMap {
		fmt.Println(k, v)
	}

	// var meuMap2 map[string]any // nil

	// meuMap2["chave"] = "valor" // panic: assignment to entry in nil map -> pq o map é nil, já q eu não inicializei ele

	// pra criar um map vazio ao invés de nil, inicializo com chaves ou make

	meuMap2 := make(map[string]any)

	meuMap2["chave"] = 5

	// meuValor := meuMap2["chaave"]          // nil
	// meuValor2 := meuMap2["chave"].(string) // panic: interface conversion: interface {} is int, not string
	// fmt.Println(meuValor)
	// fmt.Println(meuValor2)

	meuValor3, ok := meuMap2["chaave"].(string)
	if ok {
		fmt.Println(meuValor3)
	} else {
		fmt.Println("chave inexistente")
	}

}
