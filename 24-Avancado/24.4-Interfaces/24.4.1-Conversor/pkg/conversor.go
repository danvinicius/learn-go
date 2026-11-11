package conversor

import (
	"fmt"
)

type (
	Conversor interface {
		Converter(dados string) (resultado map[string]any, erro error)
		Formato() string
	}

	Gerenciador struct {
		conversores map[string]Conversor
	}
)

func NovoGerenciador() *Gerenciador {
	return &Gerenciador{
		make(map[string]Conversor),
	}
}

func (g *Gerenciador) RegistrarConversor(c Conversor) {
	g.conversores[c.Formato()] = c
}

func (g *Gerenciador) Processar(formato, dados string) {
	conversor, ok := g.conversores[formato]
	if !ok {
		fmt.Printf("Nenhum conversor %s registrado.\n", formato)
		return
	}

	resultado, erro := conversor.Converter(dados)
	if erro != nil {
		fmt.Printf("Erro ao converter formato %s.\n", formato)
		return
	}

	fmt.Printf("Resultado da conversão:\n formato: %s\n dados: %v\n", formato, resultado)
}
