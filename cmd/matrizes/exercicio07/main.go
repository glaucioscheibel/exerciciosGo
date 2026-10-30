package main

import "fmt"

func main() {
	var ordem int
	fmt.Print("Ordem da matriz identidade: ")
	fmt.Scan(&ordem)
	id := make([][]int, ordem)
	for i := range id {
		id[i] = make([]int, ordem)
	}
	for i := 0; i < ordem; i++ {
		id[i][i] = 1
	}
	for _, linha := range id {
		fmt.Printf("%v\n", linha)
	}
}
