package main

import "fmt"

func main() {
	var cont1, cont2, brancos, nulos, voto int
	for voto != -1 {
		fmt.Scan(&voto)
		switch voto {
		case 1:
			cont1++
		case 2:
			cont2++
		case 0:
			brancos++
		default:
			nulos++
		}
	}
	fmt.Println(cont1)
	fmt.Println(cont2)
	fmt.Println(brancos)
	fmt.Println(nulos)
}
