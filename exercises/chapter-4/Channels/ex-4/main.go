package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch:=make(chan int)
	const numWorkers = 3

	for i:=0;i<numWorkers;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data:= <-ch
			data*=2
			fmt.Println(data)
		}()
	}
	
	for i:=1;i<=numWorkers;i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}
