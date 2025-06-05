package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {
//se liga nesse metodos abaixo são revolucionarios se a gente para pra pensar
//a funçõa pode retornar mais de 1 coisa pqp muito boa essa linguagem 

	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 || trab2
	comprarSorvete := trab1 || trab2

	return comprarTV50, comprarTV32, comprarSorvete

	/*	if (trab1 && trab2){
			fmt.Println("Televisão 50 e sorvete")
		} if else (trab1 || trab2){
			fmt.Println("Televisão 32 e sorvete")
		} else {
			fmt.Println("não faz nada")
		}*/
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}