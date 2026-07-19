package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do main")
	auxiliar.Escrever() // funções que começam com letra minúscula só são visíveis no pacote que ela está

	erro := checkmail.ValidateFormat("viniccius774@gmail.com")
	fmt.Println(erro)
}

// go build -> gera executável
// go get -> instala pacote externo
// go run -> roda o app
// go install -> go build mas envia o arquivo pra onde está o go
// go mod tidy -> remove todas as deps q não estão sendo usadas
