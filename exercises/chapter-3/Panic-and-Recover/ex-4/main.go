package main

import (
	"fmt"
	"strconv"
)

func parsePositiveInt(s string) (int, error) {
	i,err:=strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if i <= 0 {
		return 0,fmt.Errorf("input must be a positive integer, got %d", i)
	}
	return i,err

}

func main() {
	input:="1"
	if n, err := parsePositiveInt(input); err != nil {
		fmt.Printf("Input %q error: %v\n", input, err)
	} else {
		fmt.Printf("Input %q parsed successfully: %d\n", input, n)
	}
}
// return errors for expected problems
// reserve panic for truly unexpected failures.