package main

import "fmt"

func main() {
	var a float64 = 1000
	b := 1
	var soma float64 = 0
	for b <= 50 {
		if b%2 == 0 {
			soma -= a / float64(b)
		} else {
			soma += a / float64(b)
		}
		a -= 3
		b++
	}
	fmt.Printf("Soma: %1.2f\n", soma)
}
