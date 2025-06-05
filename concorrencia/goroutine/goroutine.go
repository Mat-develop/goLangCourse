package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main(){
	//fale("Maria", "Pq vc não fala comigo?", 5)
	//fale("joão", "Pq vc não fala comigo?", 5)


	go fale("Maria", "Ei...", 10) // go meio que cria threads, doidera!!
	go fale("João", "Opa..",20) // mas ainda depende dos dados 

	time.Sleep(time.Second * 5) // ai peço pra esperar 
	fmt.Println("Fim!")
}