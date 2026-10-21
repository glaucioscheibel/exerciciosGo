package main

import "fmt"

func main() {
	var n int
	var fatorial int64 = 1
	fmt.Print("Número: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fatorial *= int64(i)
	}
	fmt.Printf("%d! = %d\n", n, fatorial)
}
