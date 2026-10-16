package main

import "fmt"

func main() {
	var valor float64
	var salario float64
	var parcelas int
	fmt.Print("Valor do empréstimo: ")
	fmt.Scan(&valor)
	fmt.Print("Quantidade de parcelas: ")
	fmt.Scan(&parcelas)
	fmt.Print("Salário: ")
	fmt.Scan(&salario)
	valorParcela := valor / float64(parcelas)
	if valorParcela >= salario*0.3 {
		fmt.Println("Empréstimo não aprovado.")
	} else {
		fmt.Println("Empréstimo aprovado.")
	}
}
