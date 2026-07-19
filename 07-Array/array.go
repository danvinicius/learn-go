package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int

	array1[2] = 900
	fmt.Println(array1)

	array2 := [10]string{"pos1", "pos2"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4}
	// array3[4] = 5 -> erro
	fmt.Println(array3)

	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array1))

	slice1 = append(slice1, 12)
	fmt.Println(slice1)

	// o slice "aponta" para um array
	slice2 := array3[1:3]
	array3[2] = 17
	fmt.Println(slice2)
}
