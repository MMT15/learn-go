package main

import (
	"fmt"
	"time"
)

func main() {
	resultCh := make(chan string)
	go func ()  {
		time.Sleep(time.Second * 3)
		resultCh <- "Operation complete!!"	
	}()
	select {
	case message:= <-resultCh:
		fmt.Println(message)
	case signal:= <-time.After(time.Second * 4):
		fmt.Println("Time out!! FAILED:", signal)
	}
}