package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5] int{1, 2, 3, 4, 5}
  // slice é uma parte de algo 
	// slice trabalha como um ponteiro 

	// slice define um pedaço de um array, mas não é um
	s2 := a2[1:3] // até mas n incluindo o indice 3
	fmt.Println(a2, s2)

	s3:= a2[:2]
	fmt.Println(s3)

	// podemos imaginar slice como : tamanho e um ponteiro para um elemento do array
	// da para fazer slice de um slice tb
	s4 := s2[:1]
	fmt.Println(s2,s4)

	// teste
	s5 := append(a1[:1], a2[:2]...)
	fmt.Println(s5)

}