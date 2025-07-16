Here are some exercises from Chapter I on **Packages**, formatted as a `README.md` file
```markdown
# Chapter 1: Packages Exercises

This section provides a set of exercises to help you solidify your understanding of Go packages, drawing from the concepts discussed in Chapter I of the course.

## Exercise 1: Basic Package Structure

Go programs are organized into **packages**. Every Go source file must belong to a package, declared at the top of the file. The `main` package is special, as it contains the `main()` function, which acts as the **entry point** for executable programs.

**Task:**
1.  Create a new directory for your Go project (e.g., `my_first_package_app`).
2.  Inside this directory, initialize a new Go module using `go mod init <your_module_path>` (e.g., `go mod init example.com/myfirstapp`).
3.  Create a file named `main.go`.
4.  Inside `main.go`, declare the package as `package main`.
5.  Implement a `main()` function that prints "Hello from my first Go package!" to the console using `fmt.Println()`.
6.  Run your program using `go run main.go`.

**Questions:**
*   What is the significance of `package main` in a Go program?
*   What is the role of the `main()` function?

## Exercise 2: Exported and Unexported Identifiers

In Go, **visibility** of identifiers (like variables or functions) is controlled by their casing.
*   Any identifier defined with an **upper case identifier** will be **exported** and visible from other packages.
*   Conversely, identifiers defined with a **lower case identifier** will **not be exported** and remain private to the package they are defined in.

**Task:**
1.  In your `my_first_package_app` project, create a new sub-directory named `utils` (e.g., `example.com/myfirstapp/utils`).
2.  Inside the `utils` directory, create a file named `string_utils.go`.
3.  Declare the package as `package utils`.
4.  In `string_utils.go`, define two functions:
    *   `Capitalize(s string) string`: This function should take a string and return its capitalized version (e.g., "hello" -> "Hello"). Ensure it's **exported**.
    *   `reverseString(s string) string`: This function should take a string and return its reversed version (e.g., "world" -> "dlrow"). Ensure it's **unexported**.
5.  In your `main.go` file, **import** your `utils` package using its full module path (e.g., `import "example.com/myfirstapp/utils"`).
6.  In the `main()` function, try to call both `utils.Capitalize("go programming")` and `utils.reverseString("language")`.
7.  Attempt to run your `main.go` file using `go run .`.

**Questions:**
*   What outcome do you observe when attempting to call `utils.reverseString` from `main.go`? Explain why this happens based on Go's export rules.
*   How can you instantly tell if a function or variable within a Go package is intended for external use or internal use?

## Exercise 3: Importing Multiple Packages and Aliasing

Go allows you to **import multiple packages**. You can also **alias** your imports to avoid naming collisions or for brevity. By convention, the package name is the last component of the import path.

**Task:**
1.  Continue using your `my_first_package_app` project with the `utils` package from the previous exercise.
2.  In `main.go`, **import both the built-in `fmt` package and your custom `utils` package**. Demonstrate importing them using a single import block.
3.  Modify the import statement for your `utils` package to use an **alias** (e.g., `import myUtil "example.com/myfirstapp/utils"`).
4.  Call the `Capitalize` function using the new alias (e.g., `myUtil.Capitalize("aliased import")`).
5.  Print a message to the console using `fmt.Println()` to confirm both imports are working correctly.

**Questions:**
*   Describe a scenario where **import aliasing** would be particularly beneficial in a Go project.
*   What is the default package name that Go would use if you import a package without providing an explicit alias?

## Exercise 4: Using External Dependencies

Go provides straightforward mechanisms to install and use **external packages** from sources like GitHub using commands such as `go install`. When a new dependency is added, the `go.mod` file is updated to define its requirements, and a `go.sum` file is created (or updated) to contain **cryptographic hashes** that ensure the **authenticity and integrity** of these module dependencies.

**Task:**
1.  In your `my_first_package_app` project, use the `go install` command to download a simple external logging package. The course references `github.com/rs/zerolog/log` as an example.
    *   Run: `go install github.com/rs/zerolog/log`
2.  Open your `main.go` file.
3.  **Import** the `log` sub-package from `github.com/rs/zerolog` (e.g., `import "github.com/rs/zerolog/log"`).
4.  In your `main()` function, use the imported `log` package to print an informational message (e.g., `log.Info().Msg("This is an external log message using zerolog!")`).
5.  Run your program using `go run .`.
6.  After running, observe your `go.mod` and `go.sum` files in the project root.

**Questions:**
*   How did the `go.mod` file change after adding the external dependency? What information does it now contain regarding `zerolog`?
*   What is the **primary purpose of the `go.sum` file**, and why is its presence crucial for projects with external dependencies?
*   If you later decide that `zerolog` is no longer needed, what command would you use to remove it and clean up your `go.mod` and `go.sum` files?
```