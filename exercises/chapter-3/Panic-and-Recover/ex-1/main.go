package main

import "fmt"

func stopNow() {
	fmt.Println("Attempting to stop...")
	panic("Critical unhealed event!")
	
}

func main() {
	stopNow()
	fmt.Println("this will not be printed")
}