package main

import (
	"fmt"
)

func loadConfig(filename string) string {
	if filename == "" {
		panic("FATAL: Configuration file path cannot be empty")
	}
	return fmt.Sprintf("Loaded config from: %s", filename)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic:", r)
		}
	}()

	fmt.Println("Starting application...")
	config := loadConfig("") 
	fmt.Println(config)    
}
