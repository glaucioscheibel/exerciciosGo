package main

import "fmt"

func main() {
	var qtdeMacas int
	var valor float64
	fmt.Print("Quantidade de maçãs: ")
	fmt.Scan(&qtdeMacas)
	if qtdeMacas < 12 {
		valor = float64(qtdeMacas) * 1.3
	} else {
		valor = float64(qtdeMacas) * 1.1
	}
	fmt.Printf("Total: R$ %1.2f \n", valor)
}
