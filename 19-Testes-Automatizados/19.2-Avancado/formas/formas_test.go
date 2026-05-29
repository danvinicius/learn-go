package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Rectangle{8, 10}
		areaEsperada := float64(80)
		areaRecebida := ret.area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida (%f) é diferente da área esperada (%f)", areaRecebida, areaEsperada)
			// diferente do Errorf, o Fatalf para de executar os demais testes caso um teste quebre
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		cir := Circle{12}
		areaEsperada := math.Pi * 12 * 12
		areaRecebida := cir.area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida (%f) é diferente da área esperada (%f)", areaRecebida, areaEsperada)
		}
	})
}
