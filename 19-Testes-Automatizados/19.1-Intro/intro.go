package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	endereco := "Avenida Paulista"
	tipoEndereco := enderecos.TipoDeEndereco(endereco)
	fmt.Println(tipoEndereco)
}
