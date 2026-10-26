package main

import "fmt"

func main() {
	vetor1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var vetor2 [10]int

	vetor2 = vetor1

	fmt.Printf("Vetor 1: %v\n", vetor1)
	fmt.Printf("Vetor 2: %v\n", vetor2)

	fmt.Println()
}
