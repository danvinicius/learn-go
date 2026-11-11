package main

import (
	"context"
	"fmt"
	"time"
)

type MyString string

const (
	id        MyString = "id"
	timestamp MyString = "timestamp"
)

func escreverCtx(nome string, ctx context.Context) {
	fmt.Printf("[%-6s] timestamp=%v id=%v\n", nome, ctx.Value(timestamp), ctx.Value(id))
}

func main() {
	ctxPai, cancelPai := context.WithCancel(
		context.WithValue(
			context.Background(),
			timestamp, time.Now().Format(time.RFC3339),
		))

	// os ctx filhos vão herdar o timestamp do ctx pai
	ctxFilho1, cancelFilho1 := context.WithCancel(context.WithValue(ctxPai, id, "context-filho-1"))
	ctxFilho2, _ := context.WithCancel(context.WithValue(ctxPai, id, "context-filho-2"))

	escreverCtx("context pai", ctxPai)
	escreverCtx("context filho 1", ctxFilho1)
	escreverCtx("context filho 2", ctxFilho2)

	// quando o ctx pai é cancelado, os ctxs filhos também são
	go func() {
		<- ctxPai.Done()
		fmt.Println("Contexto pai cancelado")
	}()
	go func() {
		<- ctxFilho1.Done()
		fmt.Println("Contexto filho 1 cancelado")
	}()
	go func() {
		<- ctxFilho2.Done()
		fmt.Println("Contexto filho 2 cancelado")
	}()

	time.Sleep(time.Second * 2)
	cancelFilho1() // cancela apenas ele msm
	time.Sleep(time.Second * 2)
	cancelPai() // cancela todos os ctx filhos (o ctxFilho1 já tava cancelado, então só cancela o ctxFilho2)
	time.Sleep(time.Second * 2)

}
