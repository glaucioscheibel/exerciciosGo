package main

import "fmt"

func main() {
	soma := 0
	var fim int
	fmt.Print("Soma dos números até qual valor? ")
	fmt.Scan(&fim)
	for i := 1; i <= fim; i++ {
		soma += i
	}
	fmt.Printf("Soma: %d \n", soma)
}
