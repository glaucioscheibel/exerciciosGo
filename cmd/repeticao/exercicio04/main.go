package main

import "fmt"

func main() {
	var tabuadas int
	fmt.Print("Tabuadas até ")
	fmt.Scan(&tabuadas)
	for tabuada := 1; tabuada <= tabuadas; tabuada++ {
		fmt.Printf("Tabuada de %d\n", tabuada)
		for i := 1; i <= 10; i++ {
			fmt.Printf("%2d x %2d = %2d \n", tabuada, i, tabuada*i)
		}
		fmt.Println()
	}
}
