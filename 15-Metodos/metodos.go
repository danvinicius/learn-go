package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) salvar() {
	fmt.Printf("Salvando o usuário %s no banco de dados\n", u.name)
}

func (u *user) fazerAniversario() {
	u.age++
}

func main() {
	u1 := user{"Daniel", 25}
	u1.salvar()
	u1.fazerAniversario()
	fmt.Println(u1)

}
