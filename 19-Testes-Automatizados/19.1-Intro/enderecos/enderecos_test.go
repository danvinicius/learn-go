package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	// t.Parallel() -> faz os testes rodarem em paralelo

	cenarios := []cenarioDeTeste{
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Rua da Lama", "Rua"},
		// {"Praça da Sé", "Tipo inválido"},
		{"AVENIDA BRASIL", "Avenida"},
	}

	for _, cenario := range cenarios {
		var retornoRecebido = TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido (%s) foi diferente do tipo esperado (%s)", retornoRecebido, cenario.retornoEsperado)
		}
	}
}

// go test ./... -> roda todos os testes das pastas abaixo

// go test -v -> (verbose) mostra todas as funções que foram executadas na suíte de testes

// go test --cover -> mostra a cobertura em % dos meus testes

// go test --coverprofile cobertura.txt -> gera o arquivo de cobertura

// go tool cover --func=cobertura.txt -> mostra a porcentagem de cobertura de cada função, de acordo com o arquivo gerado

// go tool cover --html=cobertura.txt -> gera um html detalhando exatamente as linhas da minha função que não estão sendo cobertas
