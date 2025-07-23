
---

# Go Modules Exercises

This set of exercises will help you understand how Go Modules work, drawing from the concepts introduced in the course. Go modules are a powerful feature for managing dependencies and project structure [i, 33].

## Exercise 1: Initializing Your First Module

**Objective**: Create a new Go module.

1.  **Create a new directory** for your project (e.g., `my_go_app`).
2.  Navigate into this directory in your terminal.
3.  **Initialize a new Go module** using the `go mod init` command [i, 34]. Choose a suitable module path (e.g., `github.com/yourusername/my_go_app`).
4.  **Verify** that a `go.mod` file has been created at the root of your directory [i, 33, 34]. Open it and inspect its contents. It should define the module's *module path* and also typically include the Go version [i, 35].

## Exercise 2: Adding and Managing Dependencies

**Objective**: Add an external dependency to your module and manage it.

1.  Inside your `my_go_app` directory, create a `main.go` file.
2.  Write a simple Go program that imports and uses an external package. For example, you can use `rsc.io/quote` (as seen in the context of workspaces where `stringutil` from an example is used, implying external packages can be added [i, 41]) to print a quote.
    ```go
    package main

    import (
        "fmt"
        "rsc.io/quote" // This is an external dependency
    )

    func main() {
        fmt.Println(quote.Hello())
    }
    ```
3.  **Run your program** using `go run main.go`.
    *   **Observe** the changes in your `go.mod` file after running the program. A `require` directive for `rsc.io/quote` should appear, defining its *dependency requirements* [i, 35].
    *   **Note** the creation of a `go.sum` file. This file contains the expected hashes of the content of the new modules, ensuring module authenticity and integrity [i, 35].
4.  **List all dependencies** of your module using the `go list -m all` command [i, 35]. You should see `rsc.io/quote` listed.
5.  **Remove the usage** of the imported package from your `main.go` file (e.g., delete the `import "rsc.io/quote"` statement and any code that uses `quote.Hello()`).
6.  **Clean up unused dependencies** using the `go mod tidy` command [i, 35].
7.  **Verify** that the dependency has been removed from your `go.mod` and `go.sum` files.

## Exercise 3: Understanding Module Files and `GOPATH`

**Objective**: Reinforce your understanding of module files and the historical context of `GOPATH`.

1.  Based on the course material, **explain in your own words**:
    *   What is the primary purpose of the `go.mod` file in a Go module? [i, 35]
    *   What is the `go.sum` file, and why is it important in the context of module dependencies? [i, 35]
    *   What is `GOPATH`, and what folders does it typically contain (`src`, `pkg`, `bin`)? [i, 34] How do Go modules relate to `GOPATH` and the prior Go development model? [i, 33]

## Exercise 4: Vendoring Dependencies (Optional/Advanced)

**Objective**: Learn how to vendor your module's dependencies.

1.  Ensure you have completed Exercise 2, step 2 (so your `go.mod` file has `rsc.io/quote` as a dependency).
2.  In your `my_go_app` directory, **generate a `vendor` directory** using the `go mod vendor` command [i, 36].
3.  **Inspect** the newly created `vendor` directory. You should see copies of the source code of your third-party packages (like `rsc.io/quote`) inside it [i, 36].
4.  **Explain** why vendoring might be used in a Go project [i, 36].

---