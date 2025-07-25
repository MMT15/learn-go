package main

import (
	"fmt"
	"time"
)

func main() {
	dataCh := make(chan int)
	go func ()  {
		for i:=1;i<=5;i++ {
			time.Sleep(time.Millisecond * 200)
			dataCh <- i
		}
		close(dataCh)	
	}()
	for {
		select {
		case val, ok := <-dataCh:
			if !ok {
				fmt.Println("Channel closed. Exiting.")
				return
			}
			fmt.Println("Received:", val)
		}
	}
}
