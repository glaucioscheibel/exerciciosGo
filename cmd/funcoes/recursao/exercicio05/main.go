package main

import "fmt"

func inverteCalc(palavra string, pos int) string {
	if pos == 0 {
		return "" + string(palavra[pos])
	}
	return string(palavra[pos]) + inverteCalc(palavra, pos-1)
}

func inverte(palavra string) string {
	return inverteCalc(palavra, len(palavra)-1)
}

func main() {
	var palavra string
	fmt.Print("Palavra: ")
	fmt.Scan(&palavra)
	fmt.Printf("O inverso de %s é %s\n", palavra, inverte(palavra))
}
