package main

import "sync"

func main() {
	var (
		cont = 100
		m = make(map[int]int)
		mu sync.Mutex
		wg sync.WaitGroup
	)

	wg.Add(cont)

	for i := range(cont) {
		go func() {
			defer wg.Done()
			mu.Lock() // trava a mutex caso ele esteja liberado, caso contrário espera a mutex ser destravado pela goroutine q a está usando
			m[i] = i
			mu.Unlock() // destrava a mutex
		}()
	}

	wg.Wait()
}