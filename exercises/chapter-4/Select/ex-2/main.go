package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func ()  {
		time.Sleep(time.Millisecond * 100)
		ch <- 2	
	}()
	select {
	case receive:= <-ch:
		fmt.Println(receive)
	default:
		fmt.Println("No channel operation ready")
	}
}
