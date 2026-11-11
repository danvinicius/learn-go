package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func SomaGenerica[T int | float64](arr []T) T {
	var soma T

	for _, i := range arr {
		soma += i
	}

	return soma
}

type MyInt int

type (
	_ constraints.Signed // ~int | ~int8 | ~int16 | ...
)

type Number interface {
	~int | float64 // ~int deixa eu usar qualquer tipo q se baseie em int e.g. o MyInt
}

func MultGenerica[T Number](a, b T) T {
	return a * b
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr2 := []float64{1.5, 2.3, 3.0, 4.8, 5.6}

	fmt.Println(SomaGenerica(arr))
	fmt.Println(SomaGenerica(arr2))

	x1 := 5
	x2 := 3
	y1 := 4.2
	y2 := 6.6
	fmt.Println(MultGenerica(x1, x2))
	fmt.Println(MultGenerica(y1, y2))
	// fmt.Println(MultGenerica(x1, y2)) // go não vai deixar, mesmo usando generics

	fmt.Println(MultGenerica(float64(x1), y2))

}
