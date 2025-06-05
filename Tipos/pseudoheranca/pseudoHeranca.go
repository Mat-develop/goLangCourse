package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimo, parece que está herdando mas é só composição
	turboLigado bool
}

func main() {
	f := ferrari{}

	// da pra fazer mas n dá pra referenciar direto igual na struct normal
	// tipo assim n dá
	// f1:= ferrari{"f20",0, true}
	f.nome = "F40" 
	f.velocidadeAtual = 0
	f.turboLigado = true
	fmt.Println(f)
	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)


}
