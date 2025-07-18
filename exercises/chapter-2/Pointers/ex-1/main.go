package main

import "fmt"

var num int = 42
var ptrNum *int

func main() {
	ptrNum=&num
	fmt.Println(num)
	fmt.Println(ptrNum)
	fmt.Println(*ptrNum)
}
