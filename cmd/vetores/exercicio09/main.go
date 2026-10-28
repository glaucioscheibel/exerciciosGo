package main

import "fmt"

func main() {
	a := [8]int{1, 2, 3, 4, 5, 20, 30, 50}
	b := [10]int{6, 7, 8, 9, 10, 25, 35, 55, 65, 200}
	var c []int
	ai := 0
	bi := 0
	ci := 0
	for ai < len(a) && bi < len(b) {
		if a[ai] < b[bi] {
			c = append(c, a[ai])
			ai++
		} else {
			c = append(c, b[bi])
			bi++
		}
		ci++
	}
	for ai < len(a) {
		c = append(c, a[ai])
		ai++
		ci++
	}
	for bi < len(b) {
		c = append(c, b[bi])
		bi++
		ci++
	}
	fmt.Printf("Mesclado: %v\n", c)
}
