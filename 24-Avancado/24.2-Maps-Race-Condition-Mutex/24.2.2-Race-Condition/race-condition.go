package main

import "sync"

// race condition só acontece quando 2+ goroutines acessando o mesmo dado simultaneamente
// e pelo menos uma delas está fazendo uma operação de escrita
// eu não sei quem vai chegar primeiro
func main() {
	meuMap := map[int]int{}

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for i := range 1000 {
			meuMap[i] = i
		}
		wg.Done()
	}()

	go func() {
		for i := 1000; i < 2000; i++ {
			meuMap[i] = i
		}
		wg.Done()
	}()

	wg.Wait()

	// fatal error: concurrent map writes
}
