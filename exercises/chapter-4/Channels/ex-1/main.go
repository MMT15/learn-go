package main

import "fmt"

func speak(arg string, ch chan string) {
	ch <- arg
}

func main() {
	ch:=make(chan string)
	go speak("Hello from goroutine!",ch)
	data := <-ch
	fmt.Println(data)

}
