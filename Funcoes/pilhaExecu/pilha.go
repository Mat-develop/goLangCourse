package main

import "runtime/debug"

func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main(){
	f1()
}
// pilha de execução
// f1 chama f2
// f2 chama f3
// f3 executa
// encerra f3 -> f2 -> f1 por ultimo retorna pra main ou n

// muito importante ter otimização e não estourar a pilha