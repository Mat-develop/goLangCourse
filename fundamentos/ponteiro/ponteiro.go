package main

import "fmt"

func alterarNome(nome string) {
	nome = "Bob" // Isso altera apenas a cópia local de 'nome'
}

func alterarNomeComPonteiro(nome *string) {
	*nome = "Bob" // O '*' desreferencia o ponteiro para acessar o valor original
}

func main() {
	i := 1

	// Go não tem aritmética de ponteiros
	// ponteiros nada mais é do que referencia de memoria
	var p *int = nil

	p = &i // pegando o endereço da variável i
	*p++   // pega o valor do endereço de memória e altera ele
	i++

	fmt.Println("\nEndereço:", p,"\nValor:", *p,"\nValor:", i,"\nEndereço:", &i)
	
	nome := "Alice"
	fmt.Println("Nome antes da função:", nome) // Saída: Alice

	alterarNome(nome)
	fmt.Println("Nome depois de alterarNome:", nome) // Saída: Alice (não mudou)

	alterarNomeComPonteiro(&nome) // O '&' pega o endereço de memória de 'nome'
	fmt.Println("Nome depois de alterarNomeComPonteiro:", nome) // Saída: Bob (mudou!)
}