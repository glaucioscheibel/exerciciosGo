package main

import "fmt"

func main() {
	var saldo float64
	fmt.Print("Valor depositado: ")
	fmt.Scan(&saldo)
	for mes := 1; mes <= 12; mes++ {
		saldo += saldo * 0.005
		fmt.Printf("Mês %2d saldo: %1.2f \n", mes, saldo)
	}
}
