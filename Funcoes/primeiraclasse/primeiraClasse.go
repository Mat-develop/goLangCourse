package main

import "fmt"

// modo 1 de criar
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	// modo 2 de criar
	sub := func(a, b int) int{
		return a - b
	}
	// use esses cidadões de primera classe somente quando não for algo 
	// muito importante algo para fazer coisas mais simples
	// bom que ela pode ser uma função dentro de outra função
	fmt.Println(sub(2,3))
}