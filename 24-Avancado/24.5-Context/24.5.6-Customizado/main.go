package main

import (
	"context"
	"custom-ctx/24.5.6-Customizado/custom"
	"fmt"
	"time"
)

func main() {
	ctxPai, _ := context.WithTimeout(context.Background(), time.Second*2)

	customCtx := custom.NewContext(ctxPai, map[string]any{
		"chave": "valor",
	})

	go func() {
		select {
		case <-customCtx.Done():
			fmt.Println("Custom context cancelado: ", customCtx.Err()) // nunca será executado, pois meu Done ignora o cancel do pai
		case <-time.After(time.Second * 5):
			fmt.Println("5s se passaram. Abortando...")
		}
	}()

	fmt.Println(customCtx.Value("chave"))

	time.Sleep(time.Second * 6)

}
