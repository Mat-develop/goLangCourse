package main

import "fmt"

// Interface que define um comportamento de saudação
type Saudador interface {
	Saudar(nome string) string
	Despedir() string
}

// Tipo que implementa Saudador (formal)
type Formal struct{}

func (f Formal) Saudar(nome string) string {
	return "Prezado(a) " + nome + ","
}

func (f Formal) Despedir() string {
	return "Atenciosamente."
}

// Tipo que implementa Saudador (informal)
type Informal struct{}

func (i Informal) Saudar(nome string) string {
	return "E aí, " + nome + "!"
}

func (i Informal) Despedir() string {
	return "Falou!"
}

// Função que usa a interface Saudador
func CumprimentarEDespedir(s Saudador, pessoa string) {
	fmt.Println(s.Saudar(pessoa))
	fmt.Println(s.Despedir())
}

func main() {
	formal := Formal{}
	informal := Informal{}

	CumprimentarEDespedir(formal, "Cliente")  // Passando um Formal
	CumprimentarEDespedir(informal, "Amigo") // Passando um Informal
	fmt.Println(informal.Saudar("Matheus"))
}