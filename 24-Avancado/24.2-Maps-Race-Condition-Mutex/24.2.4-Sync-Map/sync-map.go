package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		cont = 100
		m    sync.Map
		wg   sync.WaitGroup
	)

	wg.Add(cont)

	for i := range cont {
		go func() {
			defer wg.Done()
			m.Store(i, i)
		}()
	}

	wg.Wait()

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	value, ok := m.Load(5000)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("não encontrado")
	}

	/*

		Mutex vs SyncMap

		+-----------------------------------------------------------+---------------------------------------------------------------------------------------------+
		|										Mutex					                          |														SyncMap		                                  	  							    |
		+-----------------------------------------------------------+---------------------------------------------------------------------------------------------+
		|				É mais genérica e serve pra maioria dos casos			  | 						É mais eficiente pra dois casos específicos apenas  		     							      |
		+-----------------------------------------------------------+---------------------------------------------------------------------------------------------+
		|	                                            					    | Não é do msm tipo q os maps convencionais, podendo dar mais trabalho de compatibilidade     |
		+-----------------------------------------------------------+---------------------------------------------------------------------------------------------+

	*/
}
