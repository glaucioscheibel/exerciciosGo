package main

import "fmt"

func main() {
	a := [2][2]int{{2, -3}, {-1, 4}}
	var at [2][2]int
	for i := range a {
		for j := range a[i] {
			at[i][j] = a[i][j] * -1
		}
	}
	fmt.Printf("At: %v\n", at)
}
