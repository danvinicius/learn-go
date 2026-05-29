package main

import "fmt"

func main() {
	map1 := map[string]string{
		"nome": "dan",
	}

	map2 := map[string]map[string]string{
		"faculdade": {
			"curso": "SI",
		},
	}

	// fmt.Println(map1.nome) -> erro
	fmt.Println(map1["nome"])

	fmt.Println(map1)
	fmt.Println(map2)

	delete(map2, "faculdade")

	map2["endereco"] = map[string]string{
		"cidade": "SP",
	}

	fmt.Println(map2)
}
