package main

import "fmt"

func main() {
	numeros := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var pares []int
	var impares []int
	for _, n := range numeros {
		if n%2 == 0 {
			pares = append(pares, n)
		} else {
			impares = append(impares, n)
		}
	}
	fmt.Printf("Pares: %v\n", pares)
	fmt.Printf("Impares: %v\n", impares)
}
