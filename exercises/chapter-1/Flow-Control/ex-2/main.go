package main

import "fmt"

var day string = "Monday"

func main(){
	switch day {
	case "Monday":
		fmt.Println("It's a weekday! ")
		fallthrough
	case "Saturday":
		fmt.Println("It's weekend! ")
	case "Sunday":
		fmt.Println("It's weekend! ")	
	}
	x:=20
	switch {
		case x<20:
			fmt.Println("too low")
		case x>20:
			fmt.Println("too high")
		default:
			fmt.Println("Perfect")
	}
}
