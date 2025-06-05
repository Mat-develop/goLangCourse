package main

import "fmt"

// cria algo baseado no float64 mas não é um float64
// só segue a mesma formula
type nota float64

// This method is attached to the 'nota' type
// method O.O.P
// func (n nota) == this.
func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

// This function USES the 'entre' method
func notaParaconceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.9) {
		return "B"
	} else if n.entre(6.0, 6.9) {
		return "C"
	} else {
		return "Reprovado!"
	}
}

func main() {
	fmt.Println(notaParaconceito(9.8)) // Calls notaParaconceito, which uses entre
	fmt.Println(notaParaconceito(8.7)) // Calls notaParaconceito, which uses entre
	fmt.Println(notaParaconceito(5.4)) // Calls notaParaconceito, which uses entre

	// --- NEW EXAMPLE ---
	myGrade := nota(7.5) // Create a variable of type 'nota'
	// Now call the 'entre' method DIRECTLY on myGrade:
	fmt.Printf("Is my grade (%.1f) between 7.0 and 8.0? %t\n", myGrade, myGrade.entre(7.0, 8.0))
	fmt.Printf("Is my grade (%.1f) between 9.0 and 10.0? %t\n", myGrade, myGrade.entre(9.0, 10.0))
	// --- END NEW EXAMPLE ---
}