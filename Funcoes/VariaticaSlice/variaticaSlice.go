package main

import "fmt"

func printAproved(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados{
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// funciona com slice mas não funciona com array
	// espalhar assim... 
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	printAproved(aprovados...)
}

// (Spread Operator / Desempacotamento):
// assim...

//(Variadic Parameter / Parâmetro Variádico -parâmetro que recebe vários
// ...string