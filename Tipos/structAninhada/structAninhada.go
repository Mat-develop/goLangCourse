package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

// pedido n só referencia como pode criar um item na struct
type pedido struct {
	userID int
	itens  []item // slice lista * ponteiro do struct item 
	// mas o interessante que ela pode definir um item tb
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtde: 2, preco: 12.10}, //cria o item
			item{2, 1, 23.49}, // cria mais um 
			item{11, 100, 3.13}, // cria o ultimo
		}}
	fmt.Printf("Valor total do pedido R$%.2f\n", pedido.valorTotal())
	/*Então, podemos dizer que essas instâncias de item são criadas 
	e existem diretamente como elementos de uma coleção (o slice),
	 sem ter uma variável própria var meuItem = item{...} referenciando-as fora da lista.
	 Você as acessa através do slice (pedido.itens[0], pedido.itens[1], etc.) 
	 ou iterando sobre o slice (como no loop for range da função valorTotal).*/
	 fmt.Println(pedido.itens[0])
}