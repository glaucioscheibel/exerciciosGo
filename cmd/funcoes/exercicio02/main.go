package main

import "fmt"

func soma(a, b, c int) int {
	return a + b + c
}

func main() {
	a := 1
	b := 2
	c := 3
	fmt.Printf("soma de %d, %d e %d é %d\n", a, b, c, soma(a, b, c))
}
