package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	// como o retono é nomeado pouco importa a  ordem desses dois
	segundo = p2
	primeiro = p1

	return // retorno limpo
}

func main() {
	s1, s2 := 5, 6
	fmt.Println("Antes de trocar",s1,s2)
	s1, s2 = trocar(s1, s2)
	fmt.Println("Depois de trocar:",s1,s2)

	// mas lembra que o Go é bão dimais
	// e da pra fazer isso daqui
	x, y := 10, 12
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}