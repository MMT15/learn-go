package main

import (
	"errors"
	"fmt"
)

func divide(a,b float64) (float64,error) {
	if b == 0 {
		return 0,errors.New("ERROR: cannot divide by zero")
	}
	return a/b,nil
}

func main () {
	div,err:=divide(10,0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(div)

}