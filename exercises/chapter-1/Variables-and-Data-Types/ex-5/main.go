package main

import "fmt"

var integerVal int = 42
type MyString = string
type UserID int

func main(){
	floatVal:=float32(integerVal)
	fmt.Printf("%T, %T\n", integerVal,floatVal)

	var aliasVar MyString = "sure i do"
	fmt.Printf("%s - %T\n",aliasVar,aliasVar)

	var myID UserID = 15
	fmt.Println(myID, UserID(integerVal))
}
