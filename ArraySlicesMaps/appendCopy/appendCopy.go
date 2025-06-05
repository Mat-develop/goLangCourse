package main

import "fmt"

func main() {
// slice e copy só podem ser utilizados com slices e não arrays
// no copy (um slice só pode copiar um slice)

	array1 := [3] int{1,2,3}
	var slice1 []int


	// no append temos que enviar o que vamos dar append, 
	// no caso o primeiro arg seria o slice1 em si
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	// já no copy ele copia somente o que tiver espaço para copiar
	// como o tamnho é 2 ele não copia o terceiro elemento
	slice2 := make ([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)

	//mas seria uma copia mesmo??
	slice1[0] = 1
	fmt.Println(slice1, slice2)
	// s i m


}