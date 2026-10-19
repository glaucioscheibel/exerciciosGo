package main

import "fmt"

func main() {
	anacleto := 1.5
	felisberto := 1.1
	anos := 0
	for felisberto <= anacleto {
		anacleto += 0.02
		felisberto += 0.03
		anos++
	}
	fmt.Printf("Felisberto será maior que Anacleto em %d anos\n", anos)
}
