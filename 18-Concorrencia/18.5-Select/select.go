package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "msg channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "msg channel 2"
		}
	}()

	// for {
	// 	msg1 := <-canal1
	// 	fmt.Println(msg1)
	// 	msg2 := <-canal2
	// 	fmt.Println(msg2)
	// } dessa forma, eu despediço tempo pois o canal 1 precisa esperar o canal 2 receber uma mensagem para receber a próxima

	for {
		select {
		case msg1 := <-canal1:
			fmt.Println(msg1)

		case msg2 := <-canal2:
			fmt.Println(msg2)
		}
	}
}
