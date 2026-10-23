package main

import "fmt"

func main() {
	var val1 int
	var val2 int
	resultado := 0
	var operacao string
	fmt.Print("Valor1: ")
	fmt.Scan(&val1)
	fmt.Print("Valor2: ")
	fmt.Scan(&val2)
	fmt.Print("Operação: ")
	fmt.Scan(&operacao)
	operacao = string(operacao[0])
	switch operacao {
	case "+":
		resultado = val1 + val2
	case "-":
		resultado = val1 - val2
	case "*":
		resultado = val1 * val2
	case "/":
		resultado = val1 / val2
	}
	fmt.Printf("%d %s %d = %d\n", val1, operacao, val2, resultado)
}
