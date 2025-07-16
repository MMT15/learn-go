package main

import "fmt"

func main(){
	i:=1
	for i<=5 {
		fmt.Print(i," ")
		i++
	}
	fmt.Print("\n")

	for i:=1;i<=10;i++ {
		if(i>7) {
			break
		}
		if i%3==0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Print("\n")
}
