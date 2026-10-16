package main

import "fmt"

func main() {
	var ladoA int
	var ladoB int
	var ladoC int
	var msg string
	fmt.Print("Lado A: ")
	fmt.Scan(&ladoA)
	fmt.Print("Lado B: ")
	fmt.Scan(&ladoB)
	fmt.Print("Lado C: ")
	fmt.Scan(&ladoC)
	if ladoA < ladoB+ladoC && ladoB < ladoA+ladoC && ladoC < ladoA+ladoB {
		if ladoA == ladoB && ladoB == ladoC {
			msg = "Triângulo equilátero"
		} else if ladoA == ladoB || ladoB == ladoC || ladoC == ladoA {
			msg = "Triângulo isósceles"
		} else {
			msg = "Triângulo escaleno"
		}
	} else {
		msg = "Os lados não formam um triângulo"
	}
	fmt.Println(msg)
}
