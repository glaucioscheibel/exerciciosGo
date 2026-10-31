package main

import "fmt"

var movimento int

func hanoi(quantidade int, origem string, destino string, aux string) {
	if quantidade >= 1 {
		hanoi(quantidade-1, origem, aux, destino)
		movimento++
		fmt.Printf("%d: Movendo disco de %s para %s\n", movimento, origem, destino)
		hanoi(quantidade-1, aux, destino, origem)
	}
}

func main() {
	var quantidade int
	fmt.Print("Informe a quantidade de discos: ")
	fmt.Scan(&quantidade)
	hanoi(quantidade, "A", "B", "C")
}
