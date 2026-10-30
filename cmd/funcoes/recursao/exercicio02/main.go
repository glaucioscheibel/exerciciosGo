package main

import "fmt"

func fatorial(numero int) int {
	if numero <= 1 {
		return 1
	}
	return numero * fatorial(numero-1)
}

func main() {
	var numero int
	fmt.Print("Número: ")
	fmt.Scan(&numero)
	fmt.Printf("Fatorial: %d\n", fatorial(numero))
}
