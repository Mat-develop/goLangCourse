package main

import "fmt"

/*A interface define a assinatura do método (nome, parâmetros, retorno), 
mas não fornece a implementação (o código dentro da função).
 A implementação é fornecida pelos tipos concretos que 
"satisfazer" essa interface. */

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente(subentendido ou insinuado
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel){
	fmt.Println(x.toString())
}

func main (){
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)
}