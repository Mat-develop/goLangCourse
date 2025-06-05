package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota:", nota)
	} else {
		fmt.Println("Reprovado com nota:", nota)
	}

	// else n da return, somente no if ou fora do else. 
	// não usa () a n ser que vamos definir precedencia tipo
	// nota >= 7 && (true || false)
}

func main() {
	
	imprimirResultado(7.3)
	imprimirResultado(5.1)

}