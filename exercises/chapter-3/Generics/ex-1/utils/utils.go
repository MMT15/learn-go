package utils

import "fmt"

func PrintSlice[T any](slice []T) {
	for _, elem := range slice {
		fmt.Println(elem)
	}
}

