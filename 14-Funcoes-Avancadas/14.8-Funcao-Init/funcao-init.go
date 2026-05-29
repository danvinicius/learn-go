package main

import "fmt"

func init() {
	fmt.Println("Executa antes da main e pode ter um por arquivo")
}

func main() {
	fmt.Println("exec da main")
}
