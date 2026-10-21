package main

import "fmt"

func main() {
	var n int
	fmt.Print("Fibonacci até qual número? ")
	fmt.Scan(&n)
	a := 0
	b := 1
	fmt.Printf("%d, %d", a, b)
	for ok := true; ok; ok = (b < n) {
		aux := b + a
		a = b
		b = aux
		fmt.Printf(", %d", b)
	}
	fmt.Println()
}
