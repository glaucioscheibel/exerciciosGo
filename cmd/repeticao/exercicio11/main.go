package main

import "fmt"

func main() {
	var massa float64
	tempo := 0
	fmt.Print("Massa inicial: ")
	fmt.Scan(&massa)
	for massa >= 0.05 {
		tempo += 50
		massa /= 2
	}
	fmt.Printf("Levou %d segundos.\n", tempo)
}
