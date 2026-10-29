package main

import "fmt"

func main() {
	matriz := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	for _, linha := range matriz {
		for _, numero := range linha {
			fmt.Println(numero)
		}
	}
}
