package main

import "fmt"

func main() {
	a := 10
	b := 20

	aux := a
	a = b
	b = aux

	fmt.Printf("A:%d, B:%d \n", a, b)
}
