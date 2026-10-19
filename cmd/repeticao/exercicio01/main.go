package main

import "fmt"

func main() {
	i := 0
	soma := 0
	for i <= 100 {
		soma += i
		i++
	}
	fmt.Printf("Soma: R$ %d\n", soma)
}
