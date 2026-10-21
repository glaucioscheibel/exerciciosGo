package main

import (
	"fmt"
)

func main() {
	var valor float64
	var valorAPagar float64
	var tipo int
	fmt.Print("Valor da compra: ")
	fmt.Scan(&valor)
	fmt.Print("Tipo de pessoa: ")
	fmt.Scan(&tipo)
	switch tipo {
	case 1:
		valorAPagar = valor
	case 2:
		valorAPagar = valor - valor*0.1
	case 3:
		valorAPagar = valor - valor*0.05
	}
	fmt.Printf("Valor a pagar: %1.2f\n", valorAPagar)
}
