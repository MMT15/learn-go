package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrNegativeInput = errors.New("cannot have negative input")

func calculateSquareRoot(x float64) (float64,error) {
	rez:=math.Sqrt(x)
	if x < 0 {
		return rez, ErrNegativeInput
	}
	return rez,nil
}

func main() {
	rez,err:=calculateSquareRoot(9)
	if err != nil {
		switch {
		case errors.Is(err, ErrNegativeInput):
			fmt.Println(err)
		default:
			fmt.Println("i dont know")
		
		}
		return
	}
	fmt.Println(rez)
}
