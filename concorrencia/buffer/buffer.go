package main

import (
	"fmt"
	"time"
)

// Resumão dessa porra ele aguenta até 3 espaços no buffer
// se n mandar da um post ali <-ch o bagulho n roda mais pq o buffer ta cheio

func rotina(ch chan int) {

	ch <- 1
	
	ch <- 2
	ch <- 3
	<- ch
	ch <- 4
	fmt.Println("Executou!") // na teoria é pra rodar
	ch <- 5
	ch <- 6
}

func main()  {
	ch := make(chan int, 3) // chan int, buffer (tamanho do buffer aguenta 4 posições)
	go rotina(ch)
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}