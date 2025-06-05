package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão> ponteiro representado por *

func inc2(n *int) {
	// * é usado para desreferencias, ou seja, ter acesso ao valor
	*n++
}
func main() {
	n := 1
	inc1(n) // n soma o n mais sim soma a copia dele 
	fmt.Println(n)
	inc2(&n) // quando envia tem que mandar assim com o &
	// como está enviando o n de vdd e n a copia dele ele soma a var
	fmt.Println(n)
}