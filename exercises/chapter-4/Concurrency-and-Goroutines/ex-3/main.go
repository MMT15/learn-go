package main

import (
	"fmt"
	"time"
)

func greet(message string) {
	fmt.Println(message)
}

func main() {
	go greet("Hello world!")
	time.Sleep(time.Second * 1/10)

}
