package main

import "fmt"

func main() {
	const salarioMinimo = 1412.0
	var salario float64
	fmt.Print("Salário atual: ")
	fmt.Scan(&salario)
	if salario < salarioMinimo*3 {
		salario *= 1.5
	} else if salario < salarioMinimo*10 {
		salario *= 1.2
	} else if salario < salario*20 {
		salario *= 1.15
	} else {
		salario *= 1.1
	}
	fmt.Printf("Novo salário: R$ %1.2f \n", salario)
}
