package main

import "fmt"

func main() {
	var morango, macas float64
	var preco float64 = 0
	fmt.Print("Quantidade de morangos: ")
	_, _ = fmt.Scan(&morango)
	fmt.Print("Quantidade de maçãs: ")
	_, _ = fmt.Scan(&macas)
	if morango > 5 {
		preco += morango * 2.2
	} else {
		preco += morango * 2.5
	}
	if macas > 5 {
		preco += macas * 1.5
	} else {
		preco += macas * 1.8
	}
	if preco > 25 {
		preco -= preco * 0.1
	}
	fmt.Printf("Preço: %1.2f\n", preco)
}
