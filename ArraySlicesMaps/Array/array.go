package main

import "fmt"

func main() {
	// Os arrays são estruturas homogêneas (msm tipo de dados)
	// e estática (no sentido de tamanho)

	var notas[3] float64
	fmt.Println(notas) //normalmente começa com 0

	notas[0], notas[1], notas[2] = 7.5, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0
	for i:= 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Println(total)
	fmt.Printf("%.2f",media)
}