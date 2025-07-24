package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Printf("Worker %d finished\n", id)

}

func main() {
	go worker(1)
	go worker(2)
	go worker(3)
	go worker(4)
	go worker(5)
	time.Sleep(time.Second)

}
 
// The order of "starting" and "finished" messages is typically inconsistent across 
// runs because Go's goroutines execute asynchronously, and their scheduling is 
// non-deterministic. The main goroutine can also exit before other concurrently 
// launched goroutines complete, leading to varied output without explicit synchronization
