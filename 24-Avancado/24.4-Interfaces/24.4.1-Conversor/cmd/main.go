package main

import (
	conversor "github.com/danvinicius/interfaces/24.4.1-Conversor/pkg"
	"github.com/danvinicius/interfaces/24.4.1-Conversor/pkg/conversor/json"
	"github.com/danvinicius/interfaces/24.4.1-Conversor/pkg/conversor/xml"
)

func main() {
	gerenciador := conversor.NovoGerenciador()
	gerenciador.RegistrarConversor(json.ConversorJSON{})
	gerenciador.RegistrarConversor(xml.ConversorXML{})

	dadosJSON := `{
		"pessoa": {
			"nome": "dan",
			"idade": 25
		}
	}`

	dadosXML := `
		<Pessoa>
			<Nome>dan</Nome>
			<Idade>25</Idade>
		</Pessoa>
	`

	gerenciador.Processar("JSON", dadosJSON)
	gerenciador.Processar("XML", dadosXML)
}
