// função de ordem superior
// recebe uma ou + funcs
// retorna uma func como resultado

package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

// func exec recebe uma função como dois inteiros (tb
// recebe dois parametros adicionais para retornar a funcao com esses
// parametros) e retorna um inteiro
// pesquisar func map, reduce, filter
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiply, 3, 4)
	fmt.Println(resultado)
}

//Usar funções que recebem outras funções como argumentos permite:

/*Abstrair o comportamento: Você foca na lógica geral e delega os detalhes específicos.
Aumentar a reutilização de código: Uma função genérica pode ser usada com diferentes comportamentos.
Criar código mais flexível e adaptável: Você pode alterar o comportamento de uma função simplesmente passando uma função diferente como argumento.
Implementar padrões de projeto poderosos.
Facilitar a programação assíncrona e baseada em eventos.
Realizar operações complexas em coleções de dados de forma concisa.*/