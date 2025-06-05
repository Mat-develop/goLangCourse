package main

import "fmt"

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	// notação operação ternario
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	// isso n existe em GO
}

func main() {
	fmt.Println(obterResultado(6.2))
}