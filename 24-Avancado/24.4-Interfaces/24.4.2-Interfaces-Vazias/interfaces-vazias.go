package main

import "fmt"

func main() {
	var x = 25
	var y any = 15 // interface{}

	// z := x + y // invalid operation: x + y (mismatched types int and any)

	z := x + y.(int)
	fmt.Println(z)

	var a any

	a = 21
	a = "sdsdfdf"
	a = false
	a = 32.6

	switch v := a.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string: ", v)
	case float64:
		fmt.Println("float64: ", v)
	case bool:
		fmt.Println("bool: ", v)
	}

	var slice1 any = []int{1, 2, 3}
	var slice2 any = []int{1, 2, 3}

	if slice1 == slice2 {
		// PERIGO: vai compilar mas vai dar erro de runtime
	}
}
