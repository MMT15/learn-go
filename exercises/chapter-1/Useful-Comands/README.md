--------------------------------------------------------------------------------
Go Useful Commands: Hands-on Exercises
This set of exercises will introduce you to several essential Go command-line tools that help you write, maintain, and understand your Go code. These commands are integral to the Go development workflow, assisting with formatting, static analysis, environment inspection, and documentation.
Exercise 1: Formatting Code with 
Goal: To use go fmt to automatically format your Go source code according to Go's standardized style.
Concepts Covered:
• go fmt command.
• Go's enforced code formatting.
• Focusing on code logic rather than appearance.
Steps:
1. Create a New Directory and Module: Create a fresh directory for this exercise and initialize a Go module inside it.
2. Create an Unformatted Go File: Create a file named main.go with some intentionally poorly formatted Go code. Notice the inconsistent indentation and spacing.
3. Inspect the Unformatted Code: You can use cat main.go or open the file in a text editor to see its current state.
4. Run go fmt: Execute the go fmt command on your main.go file. This command will format the source code.
    ◦ The source notes that go fmt is enforced by the language, allowing developers to focus on how their code should work rather than how it should look.
5. Verify the Formatted Code: Now, inspect main.go again. You should see that it has been automatically formatted to Go's standard style.
    ◦ Observation: All spacing and indentation are now consistent.

--------------------------------------------------------------------------------
Exercise 2: Catching Mistakes with 
Goal: To use go vet to report likely mistakes in your Go packages.
Concepts Covered:
• go vet command.
• Identifying potential issues in your code.
Steps:
1. Create a New Directory and Module: Create a new directory and module for this exercise.
2. Create a Go File with a "Likely Mistake": Create a file named main.go with a common mistake that go vet is designed to catch, such as using fmt.Printf without proper format arguments.
3. Run go vet: Execute go vet to check for likely mistakes in your package.
    ◦ The source states that go vet reports likely mistakes in packages and will notify you of errors if a syntax mistake is made. While the example here is more semantic than strict syntax, go vet is designed to catch such common programming pitfalls.
4. Analyze the Output: go vet should report a warning about the format string mismatch.
    ◦ Expected Output (similar to):
5. Correct the Mistake and Re-run: Modify main.go to fix the format string, changing %s for age to %d (for integer).
6. Run go vet ./... again.
    ◦ Expected Output: (No output, indicating no issues found)

--------------------------------------------------------------------------------
Exercise 3: Inspecting Go Environment Variables with 
Goal: To use go env to inspect the Go environment variables.
Concepts Covered:
• go env command.
• Understanding key Go environment variables like GOOS, GOARCH, and GOPATH.
Steps:
1. Run go env: From any directory (it doesn't need to be a Go module), execute the go env command.
    ◦ This command simply prints all the Go environment information.
2. Identify Key Variables: Look through the output and find the following variables, understanding their purpose:
    ◦ GOOS: The operating system for which Go programs are built. (e.g., darwin for macOS, linux for Linux, windows for Windows).
    ◦ GOARCH: The processor architecture for which Go programs are built. (e.g., amd64).
    ◦ GOPATH: A variable that defines the root of your Go workspace, containing src, pkg, and bin folders. Although modules are now the default, GOPATH still plays a role in where Go looks for certain things or places compiled binaries.
    ◦ Example Output (excerpt):

--------------------------------------------------------------------------------
Exercise 4: Accessing Documentation with 
Goal: To use go doc to show documentation for a package or symbol.
Concepts Covered:
• go doc command.
• Accessing standard library and custom package documentation.
Steps:
1. Get Documentation for a Standard Library Package: Use go doc to view the documentation for the fmt package, which is part of the Go standard library.
    ◦ Observation: This will print the documentation for the entire fmt package to your terminal. The source mentions that go doc parses the source code and generates documentation in HTML format, often referenced in a project's README file.
2. Get Documentation for a Specific Function: Now, get documentation for a specific function within the fmt package, such as Println.
    ◦ Observation: This will display the signature and description of the fmt.Println function.

--------------------------------------------------------------------------------
Exercise 5: Exploring Other Commands with 
Goal: To discover other available Go commands and understand their basic purpose.
Concepts Covered:
• go help command.
• Brief overview of commands like go fix, go generate, go install, go clean, go build, and go test.
Steps:
1. Run go help: Execute go help to list all available Go commands.
    ◦ Observation: You'll see a long list of commands with brief descriptions, categorized. The source also provides a similar list.
2. Identify and Briefly Understand Other Commands: Based on the go help output and the provided source material, note the purpose of the following commands:
    ◦ go fix: Finds Go programs that use old APIs and rewrites them to use newer ones.
    ◦ go generate: Usually used for code generation.
    ◦ go install: Compiles and installs packages and dependencies. You might have seen this command in earlier discussions for adding new dependencies.
    ◦ go clean: Used for cleaning files that are generated by compilers. This command can also clear the test cache.
    ◦ go build: Compiles Go packages and their dependencies, usually producing a binary executable.
    ◦ go test: Automates testing of Go packages.

--------------------------------------------------------------------------------