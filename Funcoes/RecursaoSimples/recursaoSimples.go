package main

import "fmt"

// Essa é a função com recursão do curso
func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n - 1)
	}
}
/* 
fatorial(1) receives 1 from fatorial(0) and returns 1 * 1 = 1.
fatorial(2) receives 1 from fatorial(1) and returns 2 * 1 = 2.
fatorial(3) receives 2 from fatorial(2) and returns 3 * 2 = 6.
fatorial(4) receives 6 from fatorial(3) and returns 4 * 6 = 24.
fatorial(5) receives 24 from fatorial(4) and returns 5 * 24 = 120.  */

func main() {
	fmt.Println(fatorial(5))

	//resultado:= fatorial(-3) // ja daria ruim aqui
	//fmt.Println(resultado)

	
}