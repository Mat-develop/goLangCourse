package main

import "fmt"

func main() {
	// comando cria um slice com 10 elementos
	// mas lembra que tem que ter abstração de um array?
	// dai o programa cria um array com pelo menos, 10 posições
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	// Nesse caso aqui abaixo o make cria um slice com 10 posições
	// e um array com 20 posições 
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1,2,3,4,5,6,7,8,9,0)
	fmt.Println(s, len(s), cap(s))

	// Mesmo que ele cresça além da capacidade a google faz a boa 
	// aumenta automaticamente a capacidade do slice 

}