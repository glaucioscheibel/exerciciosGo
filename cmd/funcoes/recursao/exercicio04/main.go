package main

import (
	"errors"
	"fmt"
)

func fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("O número deve ser positivo")
	}
	return fibCalc(n), nil
}

func fibCalc(n int) int {
	if n <= 1 {
		return n
	}
	return fibCalc(n-1) + fibCalc(n-2)
}

func main() {
	var n int
	fmt.Print("N: ")
	fmt.Scan(&n)
	fib, err := fibonacci(n)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("O número na posição %d é %d\n", n, fib)
	}
}
