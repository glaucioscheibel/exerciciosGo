package main

import "fmt"

func soma(numero int) int {
	if numero <= 1 {
		return numero
	}
	return numero + soma(numero-1)
}

func main() {
	var numero int
	fmt.Print("Número: ")
	fmt.Scan(&numero)
	fmt.Printf("Somatória: %d\n", soma(numero))
}
