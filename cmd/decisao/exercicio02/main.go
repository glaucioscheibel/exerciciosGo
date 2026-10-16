package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var num3 int
	fmt.Print("Número 1: ")
	fmt.Scan(&num1)
	fmt.Print("Número 2: ")
	fmt.Scan(&num2)
	fmt.Print("Número 3: ")
	fmt.Scan(&num3)
	if num1 > num2 {
		if num1 > num3 {
			fmt.Printf("Maior: %d \n", num1)
		} else {
			fmt.Printf("Maior: %d \n", num3)
		}
	} else {
		if num2 > num3 {
			fmt.Printf("Maior: %d \n", num2)
		} else {
			fmt.Printf("Maior: %d \n", num3)
		}
	}
}
