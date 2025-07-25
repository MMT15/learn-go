package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "Message from channel 1"
	} ()

	go func() {
		time.Sleep(time.Millisecond * 500)
		ch2 <- "Message from channel 2"
	} ()
	
	select {
	case receive := <-ch1:
		fmt.Println(receive)
	case receive := <-ch2:
		fmt.Println(receive)
	}
}
