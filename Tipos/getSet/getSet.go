package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNome() string {
	return p.nome
}

func (p pessoa) getSobrenome() string {
	return p.sobrenome
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNome(n string) {
	p.nome = n
}
func (p *pessoa) setNomeCompleto(nomeCompleto string){
	parts := strings.Split(nomeCompleto, " ")
	p.nome = parts[0]
	p.sobrenome = parts[1]
}
func main() {
	p1 := pessoa{"Pedro", "Silva"}
	p2 := pessoa{"Matheus", "Soares"}

	fmt.Println(p1.getNome())

	p2.setNome("Ricardo")
	fmt.Println(p2.getNomeCompleto())

	p1.setNomeCompleto("Crocodilo Bombardillo")
	fmt.Println(p1.getNomeCompleto())
	fmt.Println(p1.getSobrenome())

}