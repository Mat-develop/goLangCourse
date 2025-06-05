package main

import (
	titulofunc "cod3r/Desktop/curso-golang/concorrencia/tituloFunc"
	"fmt"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		titulofunc.Titulo("https://www.cod3r.com.br", "https://www.goggle.com.br"),
		titulofunc.Titulo("https://www.amazon.com", "https://www.youtube.com.br"),
	)
	fmt.Println(<-c,"|", <-c)
	fmt.Println(<-c,"|", <-c)
}