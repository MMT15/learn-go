package main 

import "fmt"

type Rectangle struct {
	Length float64
	Width float64
}

func (r Rectangle) Area() float64 {
	area:=r.Length*r.Width
	return area
}

func (s Rectangle) Scale(factor float64) {
	s.Length*=factor
	s.Width*=factor
}

func main() {
	myRect:=Rectangle{}
	myRect.Length=10.0
	myRect.Width=5.0
	fmt.Println(myRect)
	fmt.Println(myRect.Area())
	myRect.Scale(2.0)
	fmt.Println(myRect)

}	
// the dimension of myRect did not change because when we call Scale method we make a
// copy of the Rectangle instance instead of the original instance. 
 