package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int32

func increment() {
	for range 1000 {
		atomic.AddInt32(&count, 1)
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
