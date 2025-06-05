package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	//esse for usa o i acima
	for i <= 10 { //não tem while em Go
		fmt.Println(i)
		i++
	}
	// essa já usa um outro i utilizado a variable shadowing
	// permitindo criar uma var igual somente dentro do escopo do for 
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i:=1; i<=10; i++ {
		if i%2 == 0 {
			fmt.Println("número par: ", i)
		} else {
			fmt.Println("número impar:", i)
		}
	}

	for {
		//laço infinito 
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
		// ctrl + alt + m para de executar
	}
}