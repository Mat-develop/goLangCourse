package main

import (
	"fmt"
	"time"
)

// Channel é a forma de comunicação entre go routines
// Channel é um tipo

func doisTresQuatroVezes(base int, c chan int){
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal 

	time.Sleep(time.Second)
	c <- 3 * base // enviando dados para o canal 

	time.Sleep(time.Second)
	c <- 4 * base // enviando dados para o canal 

	}	

	func main(){
		c := make(chan int)
		go doisTresQuatroVezes(2, c)
		fmt.Println("A")

		a, b := <- c, <- c // recebendo os dados do canal
		fmt.Println(a, b)
		fmt.Println("B")

		fmt.Println(<- c)
		fmt.Println("|The end|")
		// se tentar receber dados e n tiver mais nada para enviar causa deadlock
	}