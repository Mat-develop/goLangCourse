package main

// Map key, value

import "fmt"

func main() {

	// Não da pra fazer map igual embaixo
	// pois o Maps precisam ser inicializados
	//	var aprovados map[int] string

	//Use o literal de map quando você já conhece os pares chave-valor 
	// no momento da criação do map. 
	// Isso pode tornar o código mais conciso para inicializações simples.
	
	// aprovados := map[int]string{
	//	12345335: "Maria",
	//	12289981: "Pedro",
	//	89981891: "Ana",
	//	}


	//Use make quando você precisa criar um map vazio e 
	// adicionar elementos posteriormente, 
	// ou quando você quer especificar uma capacidade inicial.
	aprovados := make(map[int]string)

	aprovados[12345335] = "Maria"
	aprovados[12289981] = "Pedro"
	aprovados[89981891] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados{
		fmt.Println("cpf",cpf,"nome", nome)
	}

	fmt.Println(aprovados[89981891])
	delete(aprovados, 89981891)
	fmt.Println(aprovados)
}