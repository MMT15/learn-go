package main

import "fmt"

type Book struct{
	Title string
	Author string
}

func main() {
	var myBook = Book{Title: "Go programming", Author: "John Doe"}
	ptrBook:=&myBook
	fmt.Println((*ptrBook).Title)
	(*ptrBook).Author="John Smith"
	fmt.Println(myBook)
	newBookPtr:=new(Book)
	newBookPtr.Title = "Learning Go"
	newBookPtr.Author = "Alice Wonderland"
	fmt.Println(*newBookPtr)

}