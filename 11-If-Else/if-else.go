package main

import "fmt"

func main() {
	var num int = 10

	if num > 10 {
		fmt.Println("Maior que 10")
	} else if num < 10 {
		fmt.Println("Menor que 10")
	} else {
		fmt.Println("Engual a 10")
	}

	// if init
	if outroNum := num + 6; outroNum == 16 {
		fmt.Println("Engual a 16")
	}

	// outroNum fica no escopo do if e morre quando ele termina
}
