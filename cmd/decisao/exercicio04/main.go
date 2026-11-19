package main

import "fmt"

func main() {
	var valor, salario float64
	var parcelas int
	fmt.Print("Valor do empréstimo: ")
	_, _ = fmt.Scan(&valor)
	fmt.Print("Quantidade de parcelas: ")
	_, _ = fmt.Scan(&parcelas)
	fmt.Print("Salário: ")
	_, _ = fmt.Scan(&salario)
	valorParcela := valor / float64(parcelas)
	if valorParcela >= salario*0.3 {
		fmt.Println("Empréstimo não aprovado.")
	} else {
		fmt.Println("Empréstimo aprovado.")
	}
}
