package main

import "fmt"

func main() {
	matriz := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	graus := [3]int{90, 180, 270}
	for _, r := range graus {
		n := len(matriz)
		for i := range n {
			for j := i; j < n; j++ {
				aux := matriz[i][j]
				matriz[i][j] = matriz[j][i]
				matriz[j][i] = aux
			}
		}
		for i := range n {
			for j := 0; j < n/2; j++ {
				aux := matriz[i][j]
				matriz[i][j] = matriz[i][n-j-1]
				matriz[i][n-j-1] = aux
			}
		}
		fmt.Printf("Rotacionado em %d°:\n", r)
		for _, ints := range matriz {
			for _, anInt := range ints {
				fmt.Printf("%d ", anInt)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
