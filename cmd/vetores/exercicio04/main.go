package main

import "fmt"

func main() {
	var v1 [15]int
	var v2 [15]int
	vezes := 0
	for i := range v1 {
		fmt.Printf("v1[%d]: ", i)
		fmt.Scan(&v1[i])
	}
	for i := range v2 {
		fmt.Printf("v2[%d]: ", i)
		fmt.Scan(&v2[i])
	}
	for i := range v1 {
		if v1[i] == v2[i] {
			vezes++
		}
	}
	fmt.Printf("v1 e v2 possuem %d valores iguais nas mesmas posições\n", vezes)
}
