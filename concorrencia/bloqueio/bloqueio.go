package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante pq n tem buffer
	fmt.Println("Só depois que foi lido")
	/*
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante pq n tem buffer
	fmt.Println("agora n bloqueia")*/
}

func main(){
	c := make(chan int)// cannal sem buffer

	go rotina(c)
	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // esperando ler, mas ja foi consumido então da ruim
	fmt.Println("Fim")
}