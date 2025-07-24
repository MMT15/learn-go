package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	ch <- 5
}

func consumer(ch<- chan int) {
	data:= <-ch
	fmt.Println(data)
}

func main() {
	ch:=make(chan int)
	go producer(ch)
	go consumer(ch)
	time.Sleep(time.Second)
}
