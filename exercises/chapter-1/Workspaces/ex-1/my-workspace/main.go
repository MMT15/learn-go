package main

import (
	"fmt"
	"stringutil"
)

func main() {
	fmt.Println("Hello Go module!")
	fmt.Println(stringutil.ReverseString("Hello from the outside"))
}
