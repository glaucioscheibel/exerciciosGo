package main

import (
	"fmt"
	"math"
)

func main() {
	var base, expoente float64
	fmt.Print("Informe base: ")
	_, _ = fmt.Scan(&base)
	fmt.Print("Informe expoente: ")
	_, _ = fmt.Scan(&expoente)
	resultado := math.Pow(base, expoente)
	fmt.Printf("Resultado: %.2f \n", resultado)
}
