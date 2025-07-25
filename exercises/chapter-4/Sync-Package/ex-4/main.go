package main

import (
	"fmt"
	"sync"
)

var counter int
var so sync.Once
	
func performSetup() {
	fmt.Println("Performin one time...")
	counter+=1
}

func main() {
	var wg sync.WaitGroup

	wg.Add(5)
	for i:=1;i<=10;i++ {
		go func ()  {
			wg.Done()
			so.Do(performSetup)
		} ()
	}
	wg.Wait()
	fmt.Println(counter)
	
}