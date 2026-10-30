package main

import "fmt"

func soma(a ...int) int {
	resultado := 0
	for _, i := range a {
		resultado += i
	}
	return resultado
}

func main() {
	a := 1
	b := 2
	c := 3
	d := 4
	fmt.Printf("soma de %d, %d, %d e %d é %d\n", a, b, c, d, soma(a, b, c, d))
}
