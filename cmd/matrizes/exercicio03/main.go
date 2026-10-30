package main

import "fmt"

func main() {
	a := [2][4]int{{-10, 1, 4, 6}, {2, 3, 2, 8}}
	b := [2][4]int{{1, 8, 4, -1}, {0, 6, 3, -3}}
	var c [2][4]int
	for i := range a {
		for j := range a[i] {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	fmt.Printf("C: %v\n", c)
}
