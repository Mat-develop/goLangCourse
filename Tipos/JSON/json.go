package main

import (
	"encoding/json"
	"fmt"
)

// Maiusculo Public
// minisculo Private
// temos que usar `` para mapeamento par ao json
type produto struct{
	ID int `json: "id"`
	Nome string `json: "nome"`
	Preco float64 `json:"preco"`
	Tags []string `json:"tags"`
}

func main(){

	// struct para json
	p1:= produto{1, "Notebook", 1899.90, []string{"Promoção","Eletrônico"}}
	p1Json, _ := json.Marshal(p1) 	// p1Json é um slice de byte, _ error
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
	}

	/*
	principais usos e porquês do JSON:
•  Intercâmbio de Dados na Web (APIs): Este é o uso mais comum. Quando o seu navegador conversa com um servidor
 		(ou quando um servidor conversa com outro servidor), o JSON é a forma padrão de enviar e receber dados 
 		estruturados. Por exemplo:
•				Você acessa um site de previsão do tempo. O navegador envia uma requisição para um servidor. 
				O servidor processa e devolve os dados da previsão (temperatura, umidade, etc.) 
				em formato JSON, que o navegador então usa para mostrar na tela.

•				Quando um aplicativo de celular acessa informações de um serviço online (como dados de usuários,
				produtos, etc.), essa comunicação geralmente é feita via APIs que enviam e recebem dados em JSON.

•  			Configuração: Por ser fácil de ler, o JSON também é usado em alguns casos
			  para arquivos de configuração de softwares ou sistemas.
•  
				Armazenamento de Dados (em alguns bancos de dados): 
				Alguns bancos de dados NoSQL (como MongoDB) armazenam informações
			  em um formato muito similar ao JSON, chamados "documentos", porque ele permite uma estrutura flexível.

	*/