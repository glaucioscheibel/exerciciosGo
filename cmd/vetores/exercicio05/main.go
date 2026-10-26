package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{2, 4, 7, 13, 14, 15, 16}
	b := []int{1, 6, 7, 11, 13, 16, 18}

	var uniao []int
	uniao = slices.Concat(a, b)

	var interseccao []int
	for _, n := range a {
		if slices.Contains(b, n) {
			interseccao = append(interseccao, n)
		}
	}

	var diferenca []int
	for _, n := range a {
		if !slices.Contains(b, n) {
			diferenca = append(diferenca, n)
		}
	}

	fmt.Printf("A: %v\n", a)
	fmt.Printf("B: %v\n", b)
	fmt.Printf("União: %v\n", uniao)
	fmt.Printf("Interseccao: %v\n", interseccao)
	fmt.Printf("Diferença: %v\n", diferenca)
}
