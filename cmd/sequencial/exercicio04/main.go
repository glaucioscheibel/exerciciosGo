package main

import (
	"fmt"
	"math"
)

func main() {
	var raio float64
	var area float64
	fmt.Print("Informe o raio: ")
	fmt.Scan(&raio)
	area = math.Pi * math.Pow(raio, 2)
	fmt.Printf("Area: %.2f \n", area)
}
