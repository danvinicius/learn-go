package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{
		"nome": "Rex",
		"raca": "Dálmata"
	}`

	var c cachorro

	err := json.Unmarshal([]byte(cachorroEmJSON), &c)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

	cachorroEmJSON2 := `{
		"nome": "Toby",
		"raca": "Poodle"
	}`

	c2 := make(map[string]string)

	err2 := json.Unmarshal([]byte(cachorroEmJSON2), &c2)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(c2)

}
