package main

import "fmt"

func main() {
	matriz := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	somaPrimeiraColuna := 0
	produtoPrimeiraLinha := 1
	total := 0
	produtoDiagonal := 1
	for i := range len(matriz) {
		for j := range len(matriz[i]) {
			if j == 0 {
				somaPrimeiraColuna += matriz[i][j]
			}
			if i == 0 {
				produtoPrimeiraLinha *= matriz[i][j]
			}
			if i == j {
				produtoDiagonal *= matriz[i][j]
			}
			total += matriz[i][j]
		}
	}
	fmt.Printf("A soma dos elementos da primeira coluna: %d\n", somaPrimeiraColuna)
	fmt.Printf("O produto dos elementos da primeira linha: %d\n", produtoPrimeiraLinha)
	fmt.Printf("A soma de todos os elementos: %d\n", total)
	fmt.Printf("O produto da diagonal principal: %d\n", produtoDiagonal)
}
