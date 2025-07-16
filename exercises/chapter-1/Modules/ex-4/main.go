package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
}

// 4. vendoring might be used in a Go project to ensure dependency availability,
// control specific versions, enable offline work and improve reproducibility
