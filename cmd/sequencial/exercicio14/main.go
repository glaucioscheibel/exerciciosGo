package main

import "fmt"

func main() {
	var bruto float64
	fmt.Print("Salário bruto: ")
	fmt.Scanf("%f", &bruto)
	previdencia := bruto * 0.1
	ir := (bruto - previdencia) * 0.05
	liquido := bruto - previdencia - ir
	fmt.Printf("  Salário Bruto: R$ %1.2f\n", bruto)
	fmt.Printf("    Previdência: R$ %1.2f\n", previdencia)
	fmt.Printf("           I.R.: R$ %1.2f\n", ir)
	fmt.Printf("Salário Líquido: R$ %1.2f\n", liquido)
}
