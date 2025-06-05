// toda função é um closure
/*Um closure é uma função que "se lembra" e
pode acessar variáveis do escopo léxico
(o escopo onde a função foi definida) mesmo
depois que esse escopo já terminou de executar.
Em outras palavras, a função "fecha" sobre o ambiente
em que foi criada.*/

package main

import "fmt"

func closure() func(){
	x:=10
	var funcao = func(){
		fmt.Println(x)
	} // é tipo um this.x automatico
	return funcao
}


// basicamente closure fala de ordem de declaração de variavel
// se foi declarada dentro da variavel vai priorizar ela até pq 
// ela n recebe o meu x do main
func main(){
	x:=20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}