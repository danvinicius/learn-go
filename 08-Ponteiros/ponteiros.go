package main

import "fmt"

func main() {
	var variavel1 = 10
	var variavel2 = variavel1

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	variavel1++

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var variavel3 int = 100
	var ponteiroVariavel3 *int = &variavel3

	fmt.Println(ponteiroVariavel3)
	fmt.Println(*ponteiroVariavel3)

}
