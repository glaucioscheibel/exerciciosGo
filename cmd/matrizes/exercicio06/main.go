package main

import "fmt"

func main() {
	a := [2][3]int{{2, 3, 1}, {-1, 0, 2}}
	b := [3][2]int{{1, -2}, {0, 5}, {4, 1}}
	var c [2][2]int
	for i := range c {
		for j := range c[0] {
			for k := range a[i] {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	fmt.Printf("Matriz A: %v\n", a)
	fmt.Printf("Matriz B: %v\n", b)
	fmt.Printf("Matriz C: %v\n", c)
}
