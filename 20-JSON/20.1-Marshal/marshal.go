package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
	Cor   string `json:"-"` // não vai entrar no json
}

func main() {
	c := cachorro{"Rex", "Dálmata", 2, "branco"}

	cJSON, err := json.Marshal(c)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cJSON)
	fmt.Println(bytes.NewBuffer(cJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cJSON2, err := json.Marshal(c2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(cJSON2))
}
