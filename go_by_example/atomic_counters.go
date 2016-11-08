package main

import (
	"fmt"
	"runtime"
	"sync"
	// "sync/atomic"
	"time"
)

// Naive implementation using mutexes.
var uint64Mutex = &sync.Mutex{}

func myAddUint64(x *uint64, inc uint64) {
	uint64Mutex.Lock()
	*x += inc
	uint64Mutex.Unlock()
}

func myLoadUint64(x *uint64) uint64 {
	uint64Mutex.Lock()
	n := *x
	uint64Mutex.Unlock()
	return n
}

func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				// atomic.AddUint64(&ops, 1)
				myAddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	// opsFinal := atomic.LoadUint64(&ops)
	opsFinal := myLoadUint64(&ops)
	fmt.Println("Incremented:", opsFinal)
}
