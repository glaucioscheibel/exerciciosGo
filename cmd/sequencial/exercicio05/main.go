package main

import (
	"fmt"
	"math"
)

func main() {
	var base float64
	var expoente float64
	var resultado float64
	fmt.Print("Informe base: ")
	fmt.Scan(&base)
	fmt.Print("Informe expoente: ")
	fmt.Scan(&expoente)
	resultado = math.Pow(base, expoente)
	fmt.Printf("Resultado: %.2f \n", resultado)
}
