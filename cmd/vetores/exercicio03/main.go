package main

import "fmt"

func main() {
	vetor1 := [10]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	vetor2 := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	var vetor3 [10]int
	for i := range len(vetor1) {
		vetor3[i] = vetor1[i] + vetor2[i]
	}
	fmt.Printf("Vetor 1: %v\n", vetor1)
	fmt.Printf("Vetor 2: %v\n", vetor2)
	fmt.Printf("Vetor 3: %v\n", vetor3)
}
