package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	valocidadeAtual int
}

// isso aqui é interessante
// basicamente precisa passar ponteiro pra interface pois estaremos
// mexendo em uma das variaveis
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	ferrari1 := ferrari{"f40", false, 100}

	ferrari1.ligarTurbo()
	fmt.Println(ferrari1.turboLigado)

	// se for fazer com variavel lembrar de mandar endereço 
	// por conta do ponteiro
	var ferrari2 = &ferrari{"F20", false, 0}
	ferrari2.ligarTurbo()
	//ai tem que descompactar a bixona to ficando bom :)
	fmt.Println(*ferrari2)
	fmt.Println(ferrari2.modelo)

	// evitar metodos que mexem com dados ou interfaces kk

	var s esportivo = ferrari2
	fmt.Println(s)

}