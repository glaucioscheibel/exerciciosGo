package main

import "fmt"

func ehPrimo(numero int) bool {
	if numero < 2 {
		return false
	}
	for i := 2; i < numero/2; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var numero int
	fmt.Print("Número: ")
	fmt.Scan(&numero)
	var msg string
	if !ehPrimo(numero) {
		msg = "não "
	}
	fmt.Printf("O número %d %sé primo\n", numero, msg)
}
