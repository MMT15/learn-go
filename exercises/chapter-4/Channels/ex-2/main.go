package main

import "fmt"

//import "fmt"

func main() {
	ch:=make(chan int, 2)
	ch <- 10
	ch <- 20
	data1:= <-ch
	data2:= <-ch
	fmt.Println(data1)
	fmt.Println(data2)
}