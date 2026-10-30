package main

import "fmt"

func main() {
	a := [3][2]int{{-7, 8}, {4, 9}, {2, 1}}
	var at [2][3]int
	for j := range a[0] {
		for i := range a {
			at[j][i] = a[i][j]
		}
	}
	fmt.Printf("at: %v\n", at)
}
