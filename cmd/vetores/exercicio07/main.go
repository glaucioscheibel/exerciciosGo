package main

import (
	"fmt"
	"slices"
)

func main() {
	var matriculas []int
	aux := 0
	for len(matriculas) < 100 {
		fmt.Printf("Matrícula: %d: ", len(matriculas))
		fmt.Scan(&aux)
		if slices.Contains(matriculas, aux) {
			fmt.Printf("A matrícula %d já existe!\n", aux)
		} else {
			matriculas = append(matriculas, aux)
		}
	}
	fmt.Printf("Matrículas: %v\n", matriculas)
}
