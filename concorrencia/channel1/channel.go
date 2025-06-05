package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1) //tipo de canal, buffer

	ch <- 1 // enviado dados para o canal (escrita, post)
	//<- ch    // recebendo dados do canal (leitura, read (get))
	fmt.Println(<- ch) 
	
	time.Sleep(time.Second)
	ch <- 2
	fmt.Println(<- ch) 

	// o canal é um ponto de sincronismo 
	// ele espera o valor chegar por assim dizer
}