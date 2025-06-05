package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma ->", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Módulo / resto:", a%b)

	// bitwise 
	fmt.Println("E:", a&b) 	// 11 & 10 = 10
	fmt.Println("OU:", a|b) 	// 11 | 10 = 11
	fmt.Println("OU:", a^b)   // 11 ^ 10 = 01 
	
	// XOR 
	// 0 0 | 0s
	// 0 1 | 1
	// 1 0 | 1
	// 1 1 | 0

	c := 3.0
	d := 2.0
	
	// Outras operações usando math....
	// math só usa float64 cuidado
	fmt.Println("Maior =>", math.Max(float64(a),float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação", math.Pow(c,d)) // ao quadrado e etc.. 3¹²³..
}