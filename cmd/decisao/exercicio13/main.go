package main

import "fmt"

func main() {
	var morango float64
	var macas float64
	var preco float64 = 0
	fmt.Print("Quantidade de morangos: ")
	fmt.Scan(&morango)
	fmt.Print("Quantidade de maçãs: ")
	fmt.Scan(&macas)
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
