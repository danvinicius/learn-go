package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint16
}

func main() {
	var u user
	u.name = "Daniel"
	u.age = 25

	fmt.Println(u)

	usuario2 := user{"Cássia", 23, address{"Rua dos bobos", 0}}
	usuario3 := user{name: "Helon"}

	fmt.Println(usuario2)
	fmt.Println(usuario3)
}
