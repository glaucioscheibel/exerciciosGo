package main

import (
	"fmt"
)

func main() {
	var tipo string
	var litros float64
	var preco float64 = 0
	const alcool = 2.9
	const gasolina = 3.3
	fmt.Print("A-álcool, G-gasolina: ")
	fmt.Scan(&tipo)
	tipo = string(tipo[0])
	fmt.Print("Litros: ")
	fmt.Scan(&litros)
	if tipo == "A" {
		var desconto float64
		if litros <= 20 {
			desconto = 0.03
		} else {
			desconto = 0.05
		}
		preco = litros * (alcool - (alcool * desconto))
	} else if tipo == "G" {
		var desconto float64
		if litros <= 20 {
			desconto = 0.04
		} else {
			desconto = 0.06
		}
		preco = litros * (gasolina - (gasolina * desconto))
	} else {
		fmt.Println("Tipo desconhecido")
	}
	if preco > 0 {
		fmt.Printf("Preço: R$ %1.2f \n", preco)
	}
}
