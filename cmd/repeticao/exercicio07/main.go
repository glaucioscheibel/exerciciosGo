package main

import "fmt"

func main() {
	var numero int
	primo := true
	fmt.Print("Número: ")
	fmt.Scan(&numero)
	for i := 2; i < (numero/2)+1; i++ {
		if numero%i == 0 {
			primo = false
			break
		}
	}
	primoMsg := ""
	if !primo {
		primoMsg = "não "
	}
	fmt.Printf("O número %d %sé primo\n", numero, primoMsg)
}
