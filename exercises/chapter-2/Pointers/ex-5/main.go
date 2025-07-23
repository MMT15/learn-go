package main

import "fmt"

var a,b int =10,20
var ptrA,ptrB *int = &a,&b

func main() {
	fmt.Println(*ptrA==*ptrB)
	ptrB=&a
	fmt.Println(*ptrA==*ptrB)
	
}