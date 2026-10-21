package main

import (
	"fmt"
	"math/big"
)

func main() {
	graos := big.NewInt(1)
	total := big.NewInt(0)
	for casa := 1; casa <= 64; casa++ {
		total.Add(total, graos)
		graos.Mul(graos, big.NewInt(2))
	}
	fmt.Printf("Total de grãos: %d\n", total)

	//Versão alternativa:
	total2 := big.NewInt(0)
	total2.Exp(big.NewInt(2), big.NewInt(64), nil)
	total2.Sub(total2, big.NewInt(1))
	fmt.Printf("Total de grãos: %d\n", total2)
}
