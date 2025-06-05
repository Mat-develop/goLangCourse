package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// recebe quantidade de primos que quer que acha  e channel
// incio meio que segura variavel pra trocar saca?
func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100) // isso aqui é meio que desnecessario
				break
			}
		}
	}
	close(c)
}

func main(){
	c := make(chan int, 30)
	go primos(cap(c), c)

	// da pra aplicar um for no channel
	for primo := range c{
		fmt.Printf("%d ",primo)
	}
	fmt.Println("Fim")
}