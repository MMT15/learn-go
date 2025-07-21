package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height*r.Width
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

var items []interface{}

func main() {
	myRect:=Rectangle{10.0,5.0}
	myCirc:=Circle{3.0}
	items = append(items, 
		10,
		"hello",
		false,
		32.132,
		myCirc,
		myRect,
	)

	for _, i :=range items {
		switch i:=i.(type){
		case int:
			fmt.Println("Found an integer:",i)
		case string:
			fmt.Println("Found a string:",i)
		case bool:
			fmt.Println("Found a boolean:",i)
		case float64:
			fmt.Println("Found a float:",i)
		case Rectangle:
			fmt.Println("Found a Rectangle with area:", i.Area())
		case Circle:
			fmt.Println("Found a Circle with area:",i.Area())
		default:
			fmt.Println("unexpected:",i)
		}
	}


	value,ok:=items[1].(int)
	fmt.Println(value,ok)

}
// Empty interfaces are highly useful because they can hold a value of any type. 
// This makes them suitable for scenarios with heterogeneous or unknown data types 
// at compile time. A key example is fmt.Println, which uses an empty interface 
// to accept and display variables of various types, simplifying the printing process 
// without needing separate functions for each type
