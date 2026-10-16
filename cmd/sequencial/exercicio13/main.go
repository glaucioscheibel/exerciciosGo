package main

import (
	"fmt"
	"math"
)

func main() {
	var area float64
	var latas int
	fmt.Print("Área em m²: ")
	fmt.Scan(&area)
	latas = int(math.Ceil(area / 3 / 18))
	preco := float64(latas) * 80.0
	fmt.Printf("São necessárias %d latas. Total: R$%1.2f \n", latas, preco)
}
