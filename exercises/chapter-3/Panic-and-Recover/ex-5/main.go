package main

func main() {
	var p *int 
	*p=10

}
/*
Panic occurs because nil pointer has no valid memory address.
Dereferencing nil tries to access invalid memory.
Similarly, calling method on nil interface panics—no concrete value to call.
Both are unrecoverable developer errors in Go.
*/