package main

import "fmt"

// Essa é a função com recursão do curso
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf(" inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		factorialAnterior, _:= fatorial(n - 1)
		fmt.Println(n, " * ",factorialAnterior, "= next number below")
		return n * factorialAnterior, nil
	}
}

// Essa é a função com recursão que eu criei

func fatorial2(n int) (int, error) {
	if (n < 0){
		return 1, fmt.Errorf("número inválido: %d", n)
	} else{ 
		factorialAnterior, _:= fatorial(n - 1)
		fmt.Println(n, " * ",factorialAnterior, "= next number below")
		return n * factorialAnterior, nil}
	}


func main() {
	fmt.Println(fatorial(5))

	resultado, _ := fatorial(9)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fatorial2(7))
}