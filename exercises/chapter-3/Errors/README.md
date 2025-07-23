# Go Programming Exercises - Chapter III: Errors

Welcome to the exercises for the "Errors" section of the Go course! This chapter focuses on Go's idiomatic approach to error handling, which differs from exception-based languages [1]. You'll practice creating, returning, and handling errors, as well as understanding when to use `panic` instead.

## General Instructions

*   **Create a new Go module** for these exercises (e.g., `go mod init errors_exercises`).
*   **Create separate `.go` files** for each exercise or logically group them within one `main.go` file.
*   **Write clear and concise Go code** for each problem.
*   **Add comments** to explain your logic where necessary.
*   **Verify your solutions** by running your `main.go` file or creating test files if you feel adventurous (Note: Testing is a separate chapter [2]).

---

## Exercises

### Exercise 1: Basic Error Handling with `errors.New`

**Task:**
Create a function `divide(a, b float64) (float64, error)` that performs division.
If the divisor `b` is `0`, the function should **return an error** using `errors.New` [3]. Otherwise, it should return the result of the division and `nil` for the error [3].

**Instructions:**
1.  Define the `divide` function with the specified signature.
2.  Inside `main`, call `divide` with both valid and invalid divisors (e.g., `divide(10, 2)` and `divide(10, 0)`).
3.  **Handle the returned error** by checking if it's `nil` [3, 4]. If an error occurs, print an appropriate message; otherwise, print the result of the division. This approach is considered quite idiomatic in Go [4].

---

### Exercise 2: Adding Context to Errors with `fmt.Errorf`

**Task:**
Modify the `divide` function from Exercise 1. Instead of `errors.New`, use `fmt.Errorf` to create the error [4]. The error message should include the specific values that caused the division by zero (e.g., "cannot divide X by Y").

**Instructions:**
1.  Update the `divide` function to use `fmt.Errorf` [4].
2.  Ensure the error message provides **contextual information**, making it more informative for debugging.
3.  Test your updated function in `main` as before.

---

### Exercise 3: Using Sentinel Errors

**Task:**
Define a **sentinel error** called `ErrNegativeInput` [5]. Create a function `calculateSquareRoot(x float64) (float64, error)` that returns the square root of `x`. If `x` is negative, return `ErrNegativeInput`.

**Instructions:**
1.  Declare a package-level variable for `ErrNegativeInput` [5]. (By convention, sentinel error variables are often prefixed with `Err` [5]).
2.  Implement `calculateSquareRoot`.
3.  In `main`, call `calculateSquareRoot` with both positive and negative inputs.
4.  **Use `errors.Is`** to check specifically for `ErrNegativeInput` when handling the error [5]. Print a specific message if this error occurs, and a general message for any other potential errors. Sentinel errors are useful when you need to execute a different branch of code if a certain kind of error is encountered [5].

---

### Exercise 4: Implementing Custom Error Types

**Task:**
Define a **custom error type** (e.g., `InputError`) for the `calculateSquareRoot` function from Exercise 3. This custom error should have fields like `Value float64` and `Message string`. Make sure it implements the `error` interface by providing an `Error()` method that returns an error message as a string [6].

**Instructions:**
1.  Define the `InputError` struct with the specified fields.
2.  Implement the `Error()` method for `InputError` [6].
3.  Modify `calculateSquareRoot` to return an `*InputError` when `x` is negative.
4.  In `main`, when handling the error from `calculateSquareRoot`, **use `errors.As`** to check if the error can be unwrapped into `*InputError` [6]. If it can, access its fields (e.g., `error.Value`, `error.Message`) and print a detailed error message based on the custom error's data. `errors.As` checks if the error has a specific type, unlike `errors.Is` which examines if it's a particular error object [6].

---

### Exercise 5: When to Panic?

**Task:**
Consider a scenario where a program absolutely cannot proceed without a critical resource, like a configuration file. If this resource cannot be loaded at startup, the program should **panic** [7, 8].

**Instructions:**
1.  Create a function `loadConfig(filename string) string` that simulates loading a configuration.
2.  If the `filename` is empty (simulating a critical missing file, an **unrecoverable error** [7]), **call `panic()`** with an informative message (e.g., "FATAL: Configuration file path cannot be empty") [9].
3.  In `main`, call `loadConfig`. Observe the panic and the stack trace. The normal execution of the current goroutine stops immediately, and control returns to the caller until the program exits [9].
4.  **Optional (Advanced):** Implement a `defer` and `recover` block in `main` to catch the panic and print a graceful recovery message [7, 10]. In comments, **discuss when it's appropriate to `recover` from a panic** [7, 8]. `Panic` and `recover` can be considered similar to `try/catch` in other languages, but it's generally advised to avoid them and use errors when possible [7]. Panics are typically reserved for truly **unrecoverable errors** (like a missing critical config [7, 8]) or **developer errors** (e.g., dereferencing a `nil` pointer [8], as we discussed previously, a nil interface attempting a method call is a developer error leading to panic [Conversation]).

---