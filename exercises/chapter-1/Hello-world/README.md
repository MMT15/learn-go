Iată cele 3 exerciții pentru capitolul "Hello World", traduse în engleză și structurate ca un fișier `README.md`, bazate pe informațiile din surse și pe conversația noastră anterioară:

```markdown
# Exercises - "Hello World" Chapter

These exercises are designed to help you understand the fundamental concepts introduced in the "Hello World" chapter, including the structure of a Go program, the use of the `fmt` package, and how to run your code.

---

## Exercise 1: Your First "Hello, World!" Program

**Objective:** To write and run a simple Go program that displays "Hello, World!".

**Concepts Covered:**
*   **`package main`**: An executable program in Go must belong to the `main` package [Previous conversation reference]. The `main` package is considered the main package.
*   **`func main()`**: This is the special function that serves as the **entry point** for your application, just like in C, Java, or C#.
*   **`import "fmt"`**: The `fmt` package is part of the **Go standard library**, which is a set of core packages provided by the language. It is imported to allow you to perform formatted input/output operations [Previous conversation reference].
*   **`fmt.Println()`**: This function is included in the `fmt` package and is similar to `fmt.Print`, but it adds a new line at the end and also inserts spaces between arguments [19, Previous conversation reference].
*   **Running the code (`go run`)**: You can simply run the code using the `go run` command.

**Steps:**

1.  **Initialize a module:** Open your terminal in your desired project directory and initialize a new Go module. A module is basically a collection of Go packages.
    ```bash
    go mod init your_module_name
    ```
2.  **Create `main.go` file:** In the same directory, create a file named `main.go`.
3.  **Write the code:** Add the following code to `main.go`:
    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }
    ```
4.  **Run the program:** From the terminal, make sure you are in the directory containing `main.go` and run the program:
    ```bash
    go run main.go
    ```
    You should see "Hello, World!" displayed in the console.

---

## Exercise 2: Personalized Greeting with `fmt.Printf`

**Objective:** To use `fmt.Printf` to create a personalized message, introducing string formatting concepts.

**Concepts Covered:**
*   **`fmt.Printf`**: Also known as "Print Formatter", this function allows you to format numbers, strings, booleans, and much more.
*   **Annotation verbs (`%s`)**: These tell the function how to format the arguments. For example, `%s` is used to substitute strings.

**Steps:**

1.  **Reuse the setup:** Use the same module and `main.go` file from Exercise 1.
2.  **Declare a variable:** Inside the `main()` function, declare a string variable to hold a name.
    ```go
    // ... inside the main() function
    name := "Go Gopher"
    ```
3.  **Use `fmt.Printf`:** Replace the `fmt.Println` call with `fmt.Printf` to display a personalized greeting message:
    ```go
    package main

    import "fmt"

    func main() {
        name := "Go Gopher" // You can use any name you like here
        fmt.Printf("Hello, %s!\n", name)
    }
    ```
    **Note:** We added `\n` at the end to ensure a new line after the message, as `fmt.Printf` does not automatically add a new line.
4.  **Run the program:**
    ```bash
    go run main.go
    ```
    Check if the displayed message includes the name you specified.

---

## Exercise 3: Exploring Differences Between `fmt` Functions

**Objective:** To explore the behavioral differences between `fmt.Print`, `fmt.Println`, and `fmt.Printf` when displaying multiple arguments or values of different types.

**Concepts Covered:**
*   **`fmt.Print`**: Simply takes a string and prints it.
*   **`fmt.Println`**: Similar to `fmt.Print`, but it adds a new line at the end and also inserts spaces between arguments.
*   **`fmt.Printf`**: Allows controlled formatting with annotation verbs.

**Steps:**

1.  **Modify `main.go`:** In your `main.go` file, add the following lines inside the `main()` function, after any existing code:
    ```go
    package main

    import "fmt"

    func main() {
        // Exercise 2, you can leave or comment out this line
        // name := "Go Gopher"
        // fmt.Printf("Hello, %s!\n", name)

        // Exercise 3: Differences between fmt functions
        fmt.Print("This is ", "an example ", "with Print.\n")
        fmt.Println("This", "is", "an", "example", "with", "Println.")
        
        number := 101
        text := "message"
        fmt.Printf("This is a %s with the number %d.\n", text, number)
    }
    ```
2.  **Run and observe:**
    ```bash
    go run main.go
    ```
3.  **Analyze the output:**
    *   What do you observe regarding the spaces between arguments for `fmt.Print` versus `fmt.Println`?
    *   How do `fmt.Print` and `fmt.Println` handle newlines at the end, compared to `fmt.Printf`?
    *   How did annotation verbs (`%s`, `%d`) in `fmt.Printf` help you display different data types in a single formatted string?

These exercises should help you better understand the basic elements of a Go program, as presented in the "Hello World" chapter.
```