package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta pra gente
	// essa anotação [...] basicamente fala pra contar e colocar o tamanho
	// automaticamente
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}
	// for in range... 
	for _, num := range numeros{
		fmt.Println(num)
	}

	// único modo de n utilizar uma var é utilizando a notação _
}