// go run demo2.go
// go tool trace trace.out
package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func cpuBound(id int, iters int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := id
	for i := 0; i < iters; i++ {
		sum += i ^ (i >> 3)
		if i%500000 == 0 {
			// allow preemption points
			runtime.Gosched()
		}
	}
	_ = sum
}

func ioBlocked(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate "blocking" on sleep to show goroutine parking/unparking
	time.Sleep(300 * time.Millisecond)
	_ = id
}

func main() {
	// Encourage multiple Ps to see scheduling across threads
	runtime.GOMAXPROCS(4)

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		panic(err)
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// Mix cpu-bound and "IO"-blocked goroutines
	cpuG := 6
	ioG := 4

	wg.Add(cpuG + ioG)
	for i := 0; i < cpuG; i++ {
		it := 3_000_000 + rand.Intn(1_000_000)
		go cpuBound(i, it, &wg)
	}
	for i := 0; i < ioG; i++ {
		go ioBlocked(i, &wg)
	}

	wg.Wait()
	fmt.Println("done; trace written to trace.out")
}
