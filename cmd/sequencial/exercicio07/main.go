package main

import "fmt"

func main() {
	var c float64
	fmt.Print("Temperatura em celsius: ")
	_, _ = fmt.Scan(&c)
	f := (c * 9 / 5) + 32
	fmt.Printf("Temperatura em fahrenheit: %.2f\n", f)
}
