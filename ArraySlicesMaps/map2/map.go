package main

import "fmt"

func main() {
	funcioSalarios := map[string]float64{
		"José João":     11325.46,
		"Gabriel Silva": 15456.78,
		"Pedro Junior":  1200.00,
	}
	// adicionando
	funcioSalarios["Rafale Filho"] = 1350.00
	delete(funcioSalarios, "Inexistente")
	// se deletar algo que não existe não tem b.o 
	// o go entende 
	

	somaSalarios:=0.0
	for nome, salario := range funcioSalarios {
		fmt.Println(nome,": R$",salario)
		somaSalarios += salario
	}
	fmt.Printf("Soma de salários %.2f\n", somaSalarios)
	fmt.Println(funcioSalarios["Pedro Junior"])
}