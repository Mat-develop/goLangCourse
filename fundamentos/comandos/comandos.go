package main

import "fmt"

func main() {
	fmt.Printf("outro programa em %s!","go")
}

// ls - Listar arquivos e diretórios
// go - Listar pacotes e dependências do go 
// go help - Ajuda do go mostrando como certo comando funciona
// go env - Mostra as variáveis de ambiente do go
// go env GOROOT - Mostra o diretório de instalação do go
// go doc cmv/vet - reporta contruções suspeitas 

/* Compilando pela linha de comando
go run .nomedodocumento - compila o código fonte e gera um executável
go build .nomedodocumento -  gera um executável
depois .nomedodocuemnto */