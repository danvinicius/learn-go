package main

import (
	"fmt"
	"reflect" // implementa reflexão em tempo de execução -> essa lib não tem a performance boa / preferir usar o casting padrão
)

func main() {
	var x any = "32"

	t := reflect.TypeOf(x)
	fmt.Println(t) // string

	v := reflect.ValueOf(x)
	fmt.Println(v) // 32

	// xInt := v.Int() // ele compila mas vai dar panic
	// fmt.Println(xInt)

	canInt := v.CanInt()
	if canInt {
		xInt := v.Int() // muito perigo, se minha variável for uma string por exmeplo ele da panic
		fmt.Println(xInt)
	} else {
		fmt.Println("variável não compatível com o tipo int")

	}

}
