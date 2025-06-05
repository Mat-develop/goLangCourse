package main

import "fmt"

// Struc té um tipo de dado composto que permite
//  agrupar diferentes tipos de variáveis em uma única unidade
// Struct parece classe
type produto struct{
	nome string
	preco float64
	desconto float64
}

// receiver -> receptor 
// Método : func with receiver
// declara func (receiver quem recebe) resto da função normal
func(p produto) precoComDesconto() float64{
	return p.preco - (p.preco * p.desconto) 
	// p.preco *(1 - p.desconto) essa é mais elegante
}

func main(){
	// quase um constructor
	produto1 := produto{
		nome: "Lapis",
		preco: 1.79,
		desconto: 0.05,	
	}
	fmt.Println(produto1, produto1.precoComDesconto())
	// ou assim melhor ainda
	produto2 := produto{"Notebook", 1899.90, 0.10}
	fmt.Println(produto2.precoComDesconto())
}
