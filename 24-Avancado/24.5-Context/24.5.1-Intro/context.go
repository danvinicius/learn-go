package main

import (
	"context"
	"fmt"
)

type (
	MeuTipo1 string
	MeuTipo2 string
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, MeuTipo1("chave"), "valor 1") // se n tipar a entrada, o go reclama que pode haver possiveis colisões (chaves que sobrescrevem outras)
	ctx = context.WithValue(ctx, MeuTipo2("chave"), "valor 2")
	leitura(ctx)
}

func leitura(ctx context.Context) {
	v := ctx.Value(MeuTipo1("chave"))
	if v != nil {
		fmt.Println("Valor encontrado:", v)
		return
	}
	fmt.Println("Valor não encontrado")
}