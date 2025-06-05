package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int // ponteiro de inteiro

		fmt.Printf(`
	valor default int %v 
	valor default float %v  
	valor default bool %v  
	valor default string %q  
	valor default ponteiro %v`, a, b, c, d, e)
	// nil == null
}