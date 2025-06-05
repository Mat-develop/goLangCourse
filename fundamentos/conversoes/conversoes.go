package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	// int e float  conversões
	fmt.Println(x / float64(y)) // veja que temos que converter 
	// retira o decimal e n arredonda 
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// String cuidado (voltaremos mais embaixo)
	fmt.Println("Teste"+ string(97))

	// int para string
	fmt.Println("Teste" + strconv.Itoa(123))
	numero := fmt.Sprint(123)
	fmt.Println("Teste" + numero )

	// or 
	num := 123
	fmt.Println("Teste", num)

	// string para int 
	num2, _ := strconv.Atoi("123")
	fmt.Println(num2 - 122)

	//conversão de string para bool
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("verdadeiro")
	}
}