package main

import (
	"fmt"
	"sync"
	"time"
)

func hora() {
	fmt.Printf("Hora: %s\n", time.Now().Format("15:04:05.00000"))
}

func main() {
	var wg sync.WaitGroup
	for range 10 {
		wg.Go(func() {
			hora()
		})
	}
	wg.Wait()
}
