package main

import "fmt"

type Configuration struct {
	ServerHost string 
	serverPort int
	DebugMode bool
}

func main() {
	var configuration1 = Configuration{
		ServerHost: "Google", 
		serverPort: 5000, 
		DebugMode: true,
	}
	fmt.Println(configuration1.ServerHost, configuration1.DebugMode)
	fmt.Println(configuration1.serverPort)
}

// severPort should not be accessible because it does not 
// have an upper case, but in this case it is accesible because it is 
// in the same package
