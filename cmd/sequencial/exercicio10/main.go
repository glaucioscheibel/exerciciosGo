package main

import "fmt"

func main() {
	var anos int
	var meses int
	var dias int
	fmt.Print("Anos: ")
	fmt.Scan(&anos)
	fmt.Print("Meses: ")
	fmt.Scan(&meses)
	fmt.Print("Dias: ")
	fmt.Scan(&dias)
	dias += anos*365 + meses*30
	fmt.Printf("Você ja viveu %d dias\n", dias)
}
