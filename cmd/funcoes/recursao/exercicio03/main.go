package main

import "fmt"

func mdc(x, y int) int {
	if y == 0 {
		return x
	}
	return mdc(y, x%y)
}

func main() {
	var x, y int
	fmt.Print("X: ")
	fmt.Scan(&x)
	fmt.Print("Y: ")
	fmt.Scan(&y)
	fmt.Printf("O máximo divisor comum de %d e %d é %d\n", x, y, mdc(x, y))
}
