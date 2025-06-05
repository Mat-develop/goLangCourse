package main

import (
	"fmt"
	"math"
	"reflect"
)

func main(){
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	// u de unsigned 
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))
	
	// com sinal.. int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é ", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (flot32, float64)
	var x float32  = 49.99
	//or
	// var x = float32(49.99)
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	println("o tipo de literal 49.99 é", reflect.TypeOf(49.99))

	//boolean 
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) // Not
	fmt.Println(bo || false) // or 
	fmt.Println(bo && false) // and 

	// string
	s1 := "Olá meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1),reflect.TypeOf(s1)) // [á e é] valem por 2 char

	// string com multiplas linhas 
	s2 := `Olá meu nome 
	é 
	Leo`
	fmt.Println(s2 +"\n" + s1)

	// char?? go n tem char
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}