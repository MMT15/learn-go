package main

import "fmt"

var score float64 = 95.5
var ptrScore *float64 = &score

func main() {
	*ptrScore+=4.5
	fmt.Println(score)
}
