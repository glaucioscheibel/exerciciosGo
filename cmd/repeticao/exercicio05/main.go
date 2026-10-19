package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b int
	var soma float64 = 0
	fmt.Print("A: ")
	fmt.Scan(&a)
	fmt.Print("B: ")
	fmt.Scan(&b)
	for i := a; i <= b; i++ {
		soma += math.Pow(float64(i), 3)
	}
	fmt.Printf("Resultado: %2.0f\n", soma)
}
