package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "domingo"
	case 2:
		return "segunda-feira"
	case 3:
		return "terça-feira"
	case 4:
		return "quarta-feira"
	case 5:
		return "quinta-feira"
	case 6:
		return "sexta-feira"
	case 7:
		return "sábado"
	default:
		return "número inválido"
	}
}

func main() {
	var dia string = diaDaSemana(200)
	fmt.Println(dia)
}
