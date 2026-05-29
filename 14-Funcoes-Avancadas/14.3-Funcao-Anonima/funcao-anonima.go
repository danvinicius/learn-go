package main

import "fmt"

func main() {
	str := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s ", texto)
	}("Olá mundo")

	fmt.Println(str)
}
