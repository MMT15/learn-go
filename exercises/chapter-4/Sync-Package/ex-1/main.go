package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Worker %d finished!\n", id)
	
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(1,&wg)
	wg.Add(1)
	go worker(2,&wg)
	wg.Add(1)
	go worker(3,&wg)
	wg.Add(1)
	go worker(4,&wg)
	wg.Add(1)
	go worker(5,&wg)
	wg.Wait()
	
}