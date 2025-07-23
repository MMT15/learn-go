package main

import "fmt"

var uninitializedPtr *string


func main() {
	fmt.Println(uninitializedPtr)
	newBoolPtr:=new(bool)
	*newBoolPtr=true
	fmt.Println(*newBoolPtr)
	
}
