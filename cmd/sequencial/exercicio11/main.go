package main

import (
	"fmt"
	"math"
)

func main() {
	var potencia, largura, comprimento float64
	fmt.Print("Potencia da lâmpada: ")
	_, _ = fmt.Scan(&potencia)
	fmt.Print("Largura do cômodo: ")
	_, _ = fmt.Scan(&largura)
	fmt.Print("Comprimento do cômodo: ")
	_, _ = fmt.Scan(&comprimento)
	quantidade := int(math.Ceil(largura * comprimento * 18 / potencia))
	fmt.Printf("São necessárias %d lâmpadas.\n", quantidade)
}
