package main

import "fmt"

// o go não tem herança, pra fazer esse papel passamos o nome da struct pai como um atributo sem tipo

type person struct {
	name   string
	age    uint8
	height float32
	weight float32
}

type student struct {
	person
	course string
}

func main() {
	p1 := person{"daniel", 25, 1.71, 65}

	s1 := student{p1, "SI"}
	fmt.Println(s1.name)
}
