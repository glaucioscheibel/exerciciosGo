package main

import (
	"fmt"
	"strconv"
	"sync"
)

func contador(name string) {
	for i := range 10 {
		fmt.Printf("%s: %d\n", name, i)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Go(func() {
			contador("Contador " + strconv.Itoa(i+1))
		})
	}
	wg.Wait()
}
