package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x:=3.142

	xs := fmt.Sprint(x)//to string basicamente
	// Com o xs é basicamente transformar em String e somar as duas strings
	fmt.Println("O valor de x é" + xs)
	// mais facil com esse aqui
	fmt.Println("O valor de x é", x)
	// melhor ainda com esse..
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.999
	c := false 
	d := "opa"
	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n modo v %v %v %.2v %v %v", a, b, b, c, d)

}