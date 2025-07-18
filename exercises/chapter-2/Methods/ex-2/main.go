package main

import "fmt"

type Rectangle struct {
	Length float64
	Width float64
}

func (r *Rectangle) SetDimensions(newLength, newWidth float64) {
	r.Length=newLength
	r.Width=newWidth
}

func main() {
	anotherRect:=Rectangle{7.0,3.0}
	fmt.Println(anotherRect)
	anotherRect.SetDimensions(20.0,10.0)
	fmt.Println(anotherRect)
	
}

// the key difference between the SetDimensions method and the Scale method is that
// when we call Scale we work with a copy of the original instance that has the same
// values and when we call SetDimensions we work with a pointer to the original instance
// that change the values of the original instance