package main

import (
	"fmt"
	m "math" // da pra colocar uma referencia ao pacote por exemplo o m
)

func main() {
	const PI float64 = 3.1415 // Constante de ponto flutuante
	var raio = 3.2          // tipo float64 inferido pelo compilador

	// forma reduzida de declarar variáveis
	area := PI * m.Pow(raio, 2) // := forma reduzida de declarar variáveis
	// melhor explicação sobre := 
	// ele fala assim com base nessas funções que está somando,
	// declare o tipo dessa variavel ou constante enfim

	fmt.Printf("A área da circunferência é: %.2f", area) 
	// não podemos declarar uma variável sem usar ela no go.

	// Outra forma de declarar constantes e variáveis
	const (
		a = 1
		b = 2
	)
	var (
		c = 3	
		d = 4
	)
	e:= a+b+c+d 
	fmt.Println("\n",e)

	var z, f bool = true, false
	fmt.Println(z, f)

	g, h, i := 2, false, "epa!"
	print(g, h, i)
}
