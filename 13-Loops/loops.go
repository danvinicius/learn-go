package main

import (
	"fmt"
	// "time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		// time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		// time.Sleep(time.Second)
		fmt.Println(j)
	}

	for k := range 10 {
		// time.Sleep(time.Second)
		fmt.Println(k)
	}

	nomes := [3]string{"Pedro", "Tiago", "João"}
	for i, n := range nomes {
		fmt.Println(i, n)
	}
}
