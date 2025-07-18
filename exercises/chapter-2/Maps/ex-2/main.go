package main

import "fmt"

func main() {

	var userScore = make(map[string]int)
	userScore["Alice"]=100
	userScore["Bob"]=85		
	userScore["Charlie"]=92

	userScore["Alice"] = 105
	fmt.Println(userScore)
	delete(userScore, "Charlie")
	fmt.Println(userScore)

	for key, value := range userScore {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	// the order of elements printed during iteration over a map is not guaranteed.
	// in Go, maps are explicitly defined as unordered collections.
	// each iteration over the same map may result in a different order.

	copyOfUserScores:=userScore
	copyOfUserScores["Bob"]=90
	fmt.Println(userScore,copyOfUserScores)
	// in Go, maps are reference types. when you assign one map to another 
	// both variables point to the same underlying data. a modification through 
	// one reference affects the other.
}


