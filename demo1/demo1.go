// go run demo1.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i <= 10; i++ {

		go func(i int) {
			defer wg.Done()
			fmt.Printf("Printing value of i in goroutine  - %d\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("Hello, Welcome to Goroutines")
}
