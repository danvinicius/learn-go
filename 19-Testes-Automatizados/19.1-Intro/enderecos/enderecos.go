package enderecos

import "strings"

// TipoDeEndereco recebe um endereço e retorna o tipo de endereço
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "rodovia", "estrada"}

	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoMinusculo, " ")[0]

	for _, tipo := range tiposValidos {
		if primeiraPalavraEndereco == tipo {
			return strings.Title(tipo)
		}
	}
	return "Tipo inválido"
}
