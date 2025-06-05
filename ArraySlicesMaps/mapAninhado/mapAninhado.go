package main

import "fmt"

// map[string] declara o mapa e a chave
// o valor vai ser outro mapa como chave string e valor flutuante
func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 15456.78,
			"Guga Pereira":  8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome, salario := range funcs{
			fmt.Println(nome, ": R$",salario)
		}
	}
}