package main

import "fmt"

func initializeSystem(configPath string) {
	fmt.Println("Checking the path...")
	if configPath == "" {
		panic("FATAL: Required configuration file path is empty. Cannot proceed.")
	}
	fmt.Printf("System initialized successfully with config from: %s\n",configPath)
}

func main() {
	fmt.Println("Loading program...")
	initializeSystem("")
}
