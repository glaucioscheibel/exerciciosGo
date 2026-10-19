package main

import "fmt"

func main() {
	a := 37
	b := 38
	c := 1
	var soma float64 = 0
	for c <= 37 {
		soma += float64(a*b) / float64(c)
		a--
		b--
		c++
	}
	fmt.Printf("Resultado: %1.2f\n", soma)
}
