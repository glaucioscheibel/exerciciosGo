package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	var delta float64
	fmt.Print("A: ")
	_, _ = fmt.Scan(&a)
	if a == 0 {
		fmt.Println("A equação não é do segundo grau")
	} else {
		fmt.Print("B: ")
		_, _ = fmt.Scan(&b)
		fmt.Print("C: ")
		_, _ = fmt.Scan(&c)
		delta = math.Pow(float64(b), 2) - float64(4*a*c)
		if delta < 0 {
			fmt.Println("A equação não possui raízes reais")
		} else {
			r1 := (float64(-b) + math.Sqrt(delta)) / float64(2*a)
			if delta == 0 {
				fmt.Printf("A equação possui apenas uma raiz real: %1.2f\n", r1)
			} else {
				r2 := (float64(-b) - math.Sqrt(delta)) / float64(2*a)
				fmt.Printf("A equação possui duas raízes reais: %1.2f e %1.2f\n", r1, r2)
			}
		}
	}
}
