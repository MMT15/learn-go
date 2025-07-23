package main

import "fmt"

func main() {
	/*var userScore map[string]int
	userScore["Alice"]=100
	fmt.Println(userScore)
// The map `userScores` is declared but not initialized, so its value is nil.
// In Go, the zero value of a map is nil, and writing to a nil map causes a 
// runtime panic. To use the map, it must be initialized first using make().*/

	var userScore = make(map[string]int)
	userScore["Alice"]=100
	userScore["Bob"]=85		
	userScore["Charlie"]=92

	var productPrices = map[string]float64{
		"Laptop": 1200.0,
		"Mouse": 25.00,
		"Keyboard": 75.25,
	}

	fmt.Println(userScore["Alice"])
	
	fmt.Println(userScore["David"])
	// the value returned is 0 because the user David does not exist and have no value

	fmt.Println(productPrices["Laptop"])

	bob,ok:=userScore["Bob"]
	fmt.Println("Bob exists:",ok,", with the value ",bob)
}


