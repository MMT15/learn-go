package main

import "fmt"

var numbers []int

func main() {
	numbers = make([]int,3,5)
	fmt.Println(numbers,len(numbers),cap(numbers))
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	moreNumbers:=numbers
	moreNumbers[1]=25
	fmt.Println(numbers,moreNumbers)

}
// the modification to moreNumbers did affect numbers because they use the same reference,
// they point to the same address
