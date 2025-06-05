package main

import "fmt"

func media(nums ...float64) (int, float64) {
	total := 0.0
	x := 0
	// _, num se chamam identificadores
	// o primeiro identificador retorna o indice e o segundo o valor 
	// então é sabio ignora o indice dos elementos
	for i, num := range nums {
		total += num
		x = i
	}
	return x, total / float64(len(nums))
}

func main() {
	index, res := media(5.4, 3.2, 5.6, 0.0)
	fmt.Printf("%.2f last index %d",res, index+1)
}

// (Spread Operator / Desempacotamento):
// assim...

//(Variadic Parameter / Parâmetro Variádico -parâmetro que recebe vários
// ...string