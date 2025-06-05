package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	// rand.Newsource cria uma fonte de n aleatorios 
	// time.Now... cria uma semente baseado em nano segundos (noice)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	// r.Intn gera os numeros aleatorios de 0 a n(ou r no caso)
	// Sempre bom gerar a semente pois dai n fica repetindo a ordem da
	// aleatoriedade por assim dizer...
	return r.Intn(10)
}

func main(){

	// Pelo o que eu entendi da função abaixo é que nem o FOR
	// a gente cria uma variavel que fica disponivel somente para o bloco if
	// ai depois condição >>>>> if inicia var; condição
	if i := numeroAleatorio(); i>5 { //tb suportado no switch
		fmt.Println("Ganhou!!", i)
	} else{
		fmt.Println("Perdeu!!", i)
	}
}