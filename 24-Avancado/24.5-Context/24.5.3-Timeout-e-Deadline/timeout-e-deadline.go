package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id uint8) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("encerrando worker %d...\n", id)
			time.Sleep(time.Second)
			return
		default:
			fmt.Printf("continuando worker %d...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second * 2) // com o timeout, o cancel é chamado automaticamente dps do tempo, mas ele ainda retorna a função cancel por precaução
	defer cancel() // cancel de segurança para liberação de recursos 

	// ctx, cancel := context.WithDeadline(ctx, time.Now())
	// é a msm coisa que o timeout mas recebe uma data exata ao invés de uma duração (o timeout chama o deadline por trás (lá ele))

	go worker(ctx, 1)
	go worker(ctx, 2)

}
