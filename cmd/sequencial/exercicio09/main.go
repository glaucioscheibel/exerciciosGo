package main

import "fmt"

func main() {
	var comprimento float64
	var largura float64
	var altura float64
	fmt.Print("Comprimento: ")
	fmt.Scan(&comprimento)
	fmt.Print("Largura: ")
	fmt.Scan(&largura)
	fmt.Print("Altura: ")
	fmt.Scan(&altura)
	volume := comprimento * largura * altura
	fmt.Printf("Volume: %.2f \n", volume)
}
