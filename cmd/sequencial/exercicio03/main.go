package main

import "fmt"

func main() {
	var n1 float32
	var n2 float32
	var n3 float32

	var media float32

	print("Digite a primeira nota: ")
	fmt.Scanln(&n1)
	print("Digite a segunda nota: ")
	fmt.Scanln(&n2)
	print("Digite a terceira nota: ")
	fmt.Scanln(&n3)

	media = (n1*2 + n2*3 + n3*5) / 10

	fmt.Printf("A média das notas é: %.2f\n", media)
}
