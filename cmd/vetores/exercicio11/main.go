package main

import "fmt"

func main() {
	a := [9]int{1, 8, 2, 4, 3, 4, 2, 5, 1}
	b := [8]int{3, 3, 7, 5, 2, 3, 3, 7}
	var soma []int
	ai := len(a) - 1
	bi := len(b) - 1
	var vai1 int
	for ai >= 0 || bi >= 0 || vai1 > 0 {
		aux := 0
		if ai >= 0 && bi >= 0 {
			aux = a[ai] + b[bi] + vai1
		} else if ai >= 0 {
			aux = a[ai] + vai1
		} else if bi >= 0 {
			aux = b[bi] + vai1
		} else if vai1 > 0 {
			aux = vai1
		}
		if aux > 10 {
			aux %= 10
			vai1 = 1
		} else {
			vai1 = 0
		}
		soma = append([]int{aux}, soma...)
		ai--
		bi--
	}
	fmt.Printf("Resultado: %v\n", soma)
}
