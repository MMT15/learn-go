package main

import (
	"fmt"
	u "gomod/utils"
)
func main() {
	fmt.Println(u.Sum(1,2,3,4,5))
	fmt.Println(u.Sum(1.1,2.2,3.3,4.4,5.5))
	fmt.Println(u.Sum("Alice","John","Andrew"))
}
