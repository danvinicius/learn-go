package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]any
	mu   sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: map[string]any{},
	}
}

func (c *Cache) Get(key string) any {
	c.mu.Lock()
	defer c.mu.Unlock()
	if value, ok := c.data[key]; ok {
		return value
	}
	return nil
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	fmt.Printf("Valor %v armazenado no cache com a chave %s\n", value, key)
}

func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)

	for i := range 5 {
		wg.Go(func() { // msm coisa q a função debaixo
			cache.Set(fmt.Sprintf("chave_%d", i), i)
		})
	}

	for i := range 5 {
		time.Sleep(time.Millisecond * 10)
		wg.Add(1)
		go func() {
			defer wg.Done()
			value := cache.Get(fmt.Sprintf("chave_%d", i))
			fmt.Println(value)
		}()
	}

	wg.Wait()

}
