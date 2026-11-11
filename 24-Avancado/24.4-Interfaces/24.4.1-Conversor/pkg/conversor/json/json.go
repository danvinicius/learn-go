package json

import "encoding/json"

type ConversorJSON struct{}

func (c ConversorJSON) Converter(dados string) (map[string]any, error) {
	var resultado map[string]any
	err := json.Unmarshal([]byte(dados), &resultado)
	return resultado, err
}

func (c ConversorJSON) Formato() string {
	return "JSON"
}
