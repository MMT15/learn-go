package main

import "fmt"

func incrementValue(a *int) {
	*a+=1
}

func main() {
	var counter int = 10
	fmt.Println(counter)
	incrementValue(&counter)
	fmt.Println(counter)
	
}
