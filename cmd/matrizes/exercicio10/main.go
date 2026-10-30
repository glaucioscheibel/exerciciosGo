package main

import "fmt"

func main() {
	var ordem int
	fmt.Print("Ordem do triângulo de Pascal: ")
	fmt.Scan(&ordem)
	pascal := make([][]int, ordem)
	for i := 0; i < ordem; i++ {
		pascal[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				pascal[i][j] = 1
			} else {
				pascal[i][j] = pascal[i-1][j] + pascal[i-1][j-1]
			}
		}
	}
	for _, linha := range pascal {
		fmt.Printf("%v\n", linha)
	}
}
