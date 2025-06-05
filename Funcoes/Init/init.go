package main

import "fmt"

// init sempre é o primeiro a inicializar em go
// tem mais prioridade que o Main
// um inicializador
func init() {
	fmt.Println("Inicializando...")
}
// podemos ter um init para cada arquivo
func main(){
	fmt.Println("Main")
}