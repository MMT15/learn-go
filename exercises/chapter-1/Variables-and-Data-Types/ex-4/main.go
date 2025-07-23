package main

import "fmt"

var isRaining bool = true
var isCold bool = false
var val1,val2 int = 100,200

func main(){
	fmt.Print("isRaining && isCold (AND): ", isRaining&&isCold,
    "\nisRaining || isCold (OR): ", isRaining||isCold,
    "\n!isRaining` (NOT): ", !isRaining,
	"\n")

	fmt.Print("val1==val2: ", val1==val2,
	"\nval1!=val2: ", val1!=val2,
	"\nval1<val2: ", val1<val2,
	"\nval1>=val2: ", val1>=val2,
	"\n")
}
