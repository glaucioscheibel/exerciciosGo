package main

import "fmt"

func main() {
	var base int
	var expoente int
	potencia := 1
	fmt.Print("Base: ")
	fmt.Scan(&base)
	fmt.Print("Expoente: ")
	fmt.Scan(&expoente)
	for i := 0; i < expoente; i++ {
		potencia *= base
	}
	fmt.Printf("Resultado: %d\n", potencia)
}
