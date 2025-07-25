package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value+=1
	
}

func main() {
	var wg sync.WaitGroup
	var c Counter
	wg.Add(1000)
	for i:=0;i<1000;i++ {
		go func ()  {
			defer wg.Done()
			for j:=1;j<=100;j++ {
				c.Increment()
			}
		} ()
	}
	wg.Wait()
	fmt.Println(c.value)
	
}
