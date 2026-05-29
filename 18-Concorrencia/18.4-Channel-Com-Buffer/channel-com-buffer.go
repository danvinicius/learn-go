package main

import "fmt"

func main() {
	canal := make(chan string, 2) // então eu adiciono uma quantidade no buffer do canal para ele esperar alguém escutar essas msgs

	canal <- "Olá mundo"
	canal <- "Programando em Go" // se eu apenas enviar a mensagem ao canal sem buffer e não tiver ninguém para escutar, dá deadlock

	msg1 := <-canal
	msg2 := <-canal

	fmt.Println(msg1)
	fmt.Println(msg2)
}
