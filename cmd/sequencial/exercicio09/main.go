package main

import "fmt"

func main() {
	var comprimento, largura, altura float64
	fmt.Print("Comprimento: ")
	_, _ = fmt.Scan(&comprimento)
	fmt.Print("Largura: ")
	_, _ = fmt.Scan(&largura)
	fmt.Print("Altura: ")
	_, _ = fmt.Scan(&altura)
	volume := comprimento * largura * altura
	fmt.Printf("Volume: %.2f \n", volume)
}
