package main

import "fmt"

func main() {
	var base, expoente int
	potencia := 1
	fmt.Print("Base: ")
	_, _ = fmt.Scan(&base)
	fmt.Print("Expoente: ")
	_, _ = fmt.Scan(&expoente)
	for i := 0; i < expoente; i++ {
		potencia *= base
	}
	fmt.Printf("Resultado: %d\n", potencia)
}
