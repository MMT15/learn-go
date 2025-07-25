package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
	rwmu sync.RWMutex
}

func (c *Counter) Increment() {
	c.rwmu.Lock()
	defer c.rwmu.Unlock()
	c.value+=1
	
}

func (c *Counter) GetValue() int {
	c.rwmu.RLock()
	defer c.rwmu.RUnlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	var c Counter
	wg.Add(5)
	for i:=0;i<5;i++ {
		go func ()  {
			defer wg.Done()
			for j:=1;j<=10;j++ {
				c.Increment()
				time.Sleep(time.Millisecond * 500)
			}
		} ()
	}
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				val := c.GetValue()
				fmt.Printf("Reader %d read value: %d\n", id, val)
				time.Sleep(5 * time.Millisecond) // simulate faster reads
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Finale counter value:", c.value)
	
}
