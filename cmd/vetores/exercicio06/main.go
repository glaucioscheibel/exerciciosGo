package main

import "fmt"

func main() {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Printf("%v\n", v)
	aux := 0
	for i := 0; i < len(v)/2; i++ {
		aux = v[i]
		v[i] = v[len(v)-1-i]
		v[len(v)-1-i] = aux
	}
	fmt.Printf("%v\n", v)
}
