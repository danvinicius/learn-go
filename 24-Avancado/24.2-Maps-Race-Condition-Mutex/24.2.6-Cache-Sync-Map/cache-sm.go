package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data sync.Map
}

func NewCache() *Cache {
	return &Cache{
		data: sync.Map{},
	}
}

func (c *Cache) Get(key string) any {
	if value, ok := c.data.Load(key); ok {
		return value
	}
	return nil
}

func (c *Cache) Set(key string, value any) {
	c.data.Store(key, value)
	fmt.Printf("Valor %v armazenado no cache com a chave %s\n", value, key)
}

func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)

	for i := range 5 {
		wg.Go(func() {
			cache.Set(fmt.Sprintf("chave_%d", i), i)
		})
	}

	for i := range 5 {
		time.Sleep(time.Millisecond * 10)
		wg.Go(func() {
			value := cache.Get(fmt.Sprintf("chave_%d", i))
			fmt.Println(value)
		})
	}

	wg.Wait()
}
