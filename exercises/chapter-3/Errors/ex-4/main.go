package main

import (
	"errors"
	"fmt"
	"math"
)

type InputError struct {
	Value float64
	Message string
}

func (x *InputError) Error() string {
	return x.Message
}

func calculateSquareRoot(x float64) (float64,error) {
	rez:=math.Sqrt(x)
	if x < 0 {
		return rez, &InputError{x,"cannot have negative input"}
	}
	return rez,nil
}

func main() {
	rez,err:=calculateSquareRoot(-9)
	if err != nil {
		var negError *InputError
		switch {
		case errors.As(err,&negError):
			fmt.Println(negError.Message,", your input",negError.Value,"is invalid")
		default:
			fmt.Println("idk")
		}
		return
	}
	fmt.Println(rez)
}
