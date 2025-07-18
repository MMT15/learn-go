package main

import "fmt"

var name string = "Alice"
var age int = 30
var price float64 = 123.456


func main(){
	fmt.Println("Hello")
	fmt.Println("World")

	fmt.Print("Hello")
	fmt.Print("World\n")

	fmt.Printf("Hello, my name is %s and I am %d years old\n",name,age)
	
	fmt.Printf("%.2f\n",price)

	precent:=(6.75/9)*100
	fmt.Printf("%.2f%%\n", precent)

	variable:=fmt.Sprintf("The user is %s, with age %d",name,age)
	fmt.Println(variable)

	message:=`
	hello
	my
	name
	is 
	matei
	`
	fmt.Println(message)
}
