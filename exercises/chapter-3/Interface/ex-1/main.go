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

type Measurable interface {
	Area() float64
}

func PrintMeasurement(m Measurable) {
	fmt.Println("Aria = ", m.Area())
}

func main() {
	myRct:=Rectangle{10.0,5.0}
	myCirc:=Circle{3.0}
	PrintMeasurement(myRct)
	PrintMeasurement(myCirc)

}

// Rectangle and Circle can be passed to PrintMeasurement even though they are diffrent
// types because they both have a method Area() that returns float64 and when you call
// PrintMeausurement(something) it will know what type is the variable and will automaticaly
// know witch Area() method to call forward to do the calculus