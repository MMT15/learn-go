package main

// the significance of the main package is that it contains the main()
// function, witch acts as the entry point for executable programs

import (
	"fmt"
	"gomod/utils"
)

func main(){
	fmt.Println("Hello from my first Go package!")
	fmt.Println(utils.Capitalize("go programming"))
}

// when trying to export reverseString to main.go I observed that
// you cannot do that because it does not have the first character capitalized
// and this is based on Go's export rules

// you can tell instantly if a function or variable within a Go package is 
// intendet for external use or internal use just by looking at the first 
// character is upper case