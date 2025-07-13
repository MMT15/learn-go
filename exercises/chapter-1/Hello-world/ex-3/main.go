package main

import "fmt"

func main(){
	fmt.Print("This is ", "an example ", "with Print\n")
	fmt.Println("This", "is","an","example","with","Println",)
	text:="message"
	number:=101
	fmt.Printf("This is a %s with the number %d.\n",text,number)
}
