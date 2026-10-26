package main

import "fmt"

func main() {
	var numeros [100]int
	for i := range len(numeros) {
		fmt.Printf("Numero %d: ", i)
		fmt.Scan(&numeros[i])
	}
	var maior, menor, soma int
	for i := range len(numeros) {
		if i == 0 {
			maior = numeros[i]
			menor = numeros[i]
		} else {
			if numeros[i] > maior {
				maior = numeros[i]
			} else {
				menor = numeros[i]
			}
		}
		soma += numeros[i]
	}
	media := soma / len(numeros)
	fmt.Printf("Soma: %d\n", soma)
	fmt.Printf("Média: %d\n", media)
	fmt.Printf("Maior: %d\n", maior)
	fmt.Printf("Menor: %d\n", menor)
}
