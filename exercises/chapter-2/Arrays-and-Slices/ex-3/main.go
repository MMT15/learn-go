package main

import "fmt"

func main() {
	data:=[]string{"apple", "banana", "cherry"}
	newdata:= append(data, "date", "elderberry")
	fmt.Println(data, len(data),cap(data))
	fmt.Println(newdata,len(newdata),cap(newdata))
	souceSlice:=[]string{"one", "two", "three", "four"}
	destinationSlice := make([]string, 2, 4)
	copy(destinationSlice, souceSlice)
	fmt.Println(souceSlice, destinationSlice)
}
// the copy function copies elements up to the length of the smaller of the two slices
// and the destinationSlice had the length of 2 while the source destination had 
// the length of 4 so the copy function copied just the fist 2 elements from souceSlice
