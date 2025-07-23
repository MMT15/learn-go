package main

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func riskyOperation() {
	defer handlePanic() 
	panic("Oh no, something went wrong in riskyOperation!")
}

func main() {
	riskyOperation()

	fmt.Println("Program continues after recovery.")
}

// use panic/recover for unexpected, unrecoverable errors.
// use error returns for normal, expected failures.
