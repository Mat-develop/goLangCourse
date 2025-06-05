package main

import "fmt"

// não retorna nada, não recebe nada.
func f1() {
	fmt.Println("F1")
}
// recebe, mas não retorna
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}
// não recebe, mas retorna
func f3() string{
	return "F3"
}

// recebe e retona
func f4(p1 int, p2 int) int{
	return p1+p2
}

// função pura é uma função que não mexe em nada externo,
// ou seja recebe uns parametros mexe com eles mas não altera os dados
// fora da função

// go Special func
// quando retornar dois colocar em parenteses
func f5() (string, string){
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("oie", "lindão")

	resposta1 := f3()
	fmt.Println(resposta1)

	resposta2 := f4(4,6)
	fmt.Println(resposta2)

	// para receber mais de uma resposta da pra fazer assim:
	resposta3, resposta4 := f5()
	fmt.Println(resposta3,resposta4)

	// se caso quiseremos ignorar uma das respostas
	r1, _ := f5()
	fmt.Println(r1)
}