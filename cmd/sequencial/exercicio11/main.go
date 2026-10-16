package main

import (
	"fmt"
	"math"
)

func main() {
	var potencia float64
	var largura float64
	var comprimento float64
	fmt.Print("Potencia da lâmpada: ")
	fmt.Scan(&potencia)
	fmt.Print("Largura do cômodo: ")
	fmt.Scan(&largura)
	fmt.Print("Comprimento do cômodo: ")
	fmt.Scan(&comprimento)
	quantidade := int(math.Ceil(largura * comprimento * 18 / potencia))
	fmt.Printf("São necessárias %d lâmpadas.\n", quantidade)
}
