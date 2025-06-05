package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// podendo criar mais metodos
}

type bmw struct{}

func (b bmw) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw) fazerBaliza() {
	fmt.Println("Fazendo Baliza...")
}

func(b bmw) abrirTeto(){
	fmt.Println("abrindo teto")
}

func main() {
	var b  esportivoLuxuoso = bmw{}
	b.fazerBaliza()
	b.ligarTurbo()
	//b.abrirTeto() como a interface não tem esse metodo n podemos utilizar

	carro := bmw{}
	carro.fazerBaliza()
	carro.ligarTurbo()
	carro.abrirTeto()

	/*	Usar a variável carro (do tipo concreto bmw) 
			 é o que você faria normalmente para trabalhar diretamente com um objeto bmw.
			 Você tem acesso total a tudo que o bmw pode fazer.

			Usar a variável b (do tipo interface esportivoLuxuoso) é o que você 	
			 faz quando quer trabalhar com qualquer tipo que satisfaça essa interface,
			 sem se importar qual é o tipo concreto exato por baixo. 
			 Isso oferece flexibilidade e abstração. No seu exemplo, 
			 tanto bmw quanto qualquer outro tipo (digamos, ferrari ou porsche) 
			 que implemente ligarTurbo() e fazerBaliza() poderiam ser atribuídos 
			 a uma variável do tipo esportivoLuxuoso. */
}