package main

import (
	u "gomod/utils"
)
func main() {
	intSlice:=[]int{
		1,2,3,4,
	}
	stringSlice:=[]string{
		"Ana",
		"Maria",
		"Cornel",
	}
	floatSlice:=[]float64{
		1.23,
		2.3214,
		1231.13211,
	}
	u.PrintSlice(intSlice)
	u.PrintSlice(stringSlice)
	u.PrintSlice(floatSlice)
}
