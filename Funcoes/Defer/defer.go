package main

import "fmt"

// defer é usado para fechar algum tipo de recurso que abriu
// vai jogar pro último possivel (antes do return)
func getAprovedValue (num int) int {

	defer fmt.Println("fim!")
	// garante que isso aqui será executado no momento
	// mais tardiu possivel

	if num >= 5000 {
		fmt.Println("Valor alto..")
		return 5000
	} else {
		fmt.Println("Valor baixo...")
		return num
	}
}
func main() {
	fmt.Println(getAprovedValue(6500))
}