package main

import "fmt"

func notaParaConceito(n float64) string {
	switch{
	case 10 >= n && n>=9: 
		return "A"
	case 9 > n && n >= 8:
		return "B"
	case 8 > n && n >= 5:
		return "C"
	case 5 > n && n >= 3:
		return "D"
	case 3 > n && n >= 0:
		return "E"
	default:
		return "Invalid"
	
}
}

// legal fazer codigo assim, mas quando a necessidade de multiplas seleções
// utilizar o switch case

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.8))
}