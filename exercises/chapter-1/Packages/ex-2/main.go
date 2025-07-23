package main

import (
	"fmt"
	myUtils "gomod/utils"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println(myUtils.Capitalize("aliased import"))
	log.Info().Msg("This is an external log using zerolog")
}

// a scenario where import alias would be particularly beneficial in a GO 
// project would be when you have multiple packages with long and complicate 
// names and also when you have packages with the same name but diffrent paths

// the default package name that Go would use if you import a package without 
// providing an explicit alias is the name of the last path

//*   How did the `go.mod` file change after adding the external dependency? 
// What information does it now contain regarding `zerolog`?

// After adding zerolog, the go.mod file will be updated to include a require
// directive specifying the zerolog module and its version. This file defines
// the module's path and its dependency requirements

//*   What is the **primary purpose of the `go.sum` file**, and why is its 
// presence crucial for projects with external dependencies?

// The primary purpose of the go.sum file is to contain the expected hashes 
// of the content of new modules. Its presence is crucial because it ensures 
// the integrity and security of dependencies, acting as a guarantee that 
// downloaded modules are exactly as expected

//*   If you later decide that `zerolog` is no longer needed, what command 
// would you use to remove it and clean up your `go.mod` and `go.sum` files?

// go mod tidy