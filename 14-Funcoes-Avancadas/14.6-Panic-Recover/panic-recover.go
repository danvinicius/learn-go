package main

import "fmt"

func tryRecover() {
	if rec := recover(); rec != nil { // se tiver um panic, vai ser diferente de nil
		fmt.Println("Recover executado com sucesso")
	}
}

func naoEh42(num int) bool {
	defer tryRecover()
	if num != 42 {
		fmt.Println("Não eh 42")
		return true
	} else {
		panic("EH 42") // vai chamar todos os DEFER
	}
}

func main() {
	x := naoEh42(42)
	fmt.Println(x)
}
