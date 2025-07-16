package main

import "fmt"

var num int = 10

func main(){
	if num%2==0 {
		fmt.Printf("The number %d is even\n", num)
	} else {
		fmt.Printf("The number %d is odd\n", num)
	}
	
	if result:= num%2; result==0 {
		fmt.Printf("The number %d is even\n",num)
	} else {
		fmt.Printf("The number %d is odd\n",num)
	}
}
