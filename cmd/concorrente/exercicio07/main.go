package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count atomic.Int32

func increment() {
	for range 1000 {
		count.Add(1)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Go(func() {
		increment()
	})
	wg.Go(func() {
		increment()
	})
	wg.Wait()
	fmt.Println(count)
}
