package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	// usando var interface podemos ter uma variavel que recebe qualquer
	// tipo , doidera demais (como se fosse um object do java)
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}
	var coisa2 dinamico ="opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Sabu muito"}
	fmt.Println(coisa2)
}