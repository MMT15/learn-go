# Go Programming Exercises - Chapter III: Panic and Recover

Welcome to the exercises for the "Panic and Recover" section! This chapter delves into Go's mechanism for handling truly exceptional and unrecoverable situations, distinct from its idiomatic error handling. You'll practice using `panic` to stop program execution and `recover` to regain control.

## General Instructions

*   **Create a new Go module** for these exercises (e.g., `go mod init panic_recover_exercises`).
*   **Create separate `.go` files** for each exercise or logically group them within one `main.go` file.
*   **Write clear and concise Go code** for each problem.
*   **Add comments** to explain your logic where necessary.
*   **Verify your solutions** by running your `main.go` file.

---

## Exercises

### Exercise 1: Basic Panic

**Task:**
Create a simple function that calls `panic` unconditionally. Observe how `panic` immediately stops the normal execution of the current goroutine [1].

**Instructions:**
1.  Define a function `stopNow()`. Inside `stopNow`, print a message like "Attempting to stop..." and then call `panic("Critical unhandled event!")`.
2.  After the `panic` call in `stopNow`, try to print another message (e.g., "This will not be printed.").
3.  In your `main` function, call `stopNow()`.
4.  **Run your program** and observe the output. You should see the panic message and a stack trace, and the program will terminate abruptly [1].

---

### Exercise 2: Panic for an Unrecoverable Error

**Task:**
Simulate a scenario where a program cannot proceed without a crucial resource (e.g., a configuration file). If this resource is missing, the program should `panic` because it represents an **unrecoverable error** [2, 3].

**Instructions:**
1.  Create a function `initializeSystem(configPath string)`.
2.  Inside `initializeSystem`, simulate checking if `configPath` is valid. If `configPath` is empty, **call `panic()`** with a message like "FATAL: Required configuration file path is empty. Cannot proceed." [3].
3.  Otherwise (if `configPath` is not empty), print "System initialized successfully with config from: [path]".
4.  In your `main` function, call `initializeSystem` with an empty string and observe the panic.
5.  Then, call `initializeSystem` with a non-empty string (e.g., `"config.txt"`) to see the successful path.

---

### Exercise 3: Recovering from a Panic

**Task:**
Use `defer` and `recover` to gracefully handle a panic that occurs within a function. This demonstrates how to regain control of a panicking program [4].

**Instructions:**
1.  Define a function `riskyOperation()`. Inside this function, intentionally cause a `panic` (e.g., `panic("Oh no, something went wrong in riskyOperation!")`).
2.  In your `main` function, use a `defer` statement to call an anonymous function or a separate helper function (e.g., `handlePanic()`).
3.  Inside the deferred function, **call `recover()`** [4].
4.  If `recover()` returns a non-`nil` value (meaning a panic occurred), print a message indicating that the panic was recovered, along with the panic value (e.g., "Recovered from panic: [panic message]"). This allows the rest of the `main` function to continue execution [2].
5.  Place the call to `riskyOperation()` *before* a print statement in `main` that should run only if the panic is recovered (e.g., "Program continues after recovery.").
6.  **Discuss in comments** within your code when it's **appropriate to `recover` from a panic** [2]. Generally, `panic` and `recover` are similar to `try/catch` in other languages but should be avoided in favor of explicit `error` returns for recoverable conditions [2].

---

### Exercise 4: When to Use Errors (Not Panic)

**Task:**
Create a function that performs an operation that *could* fail (e.g., parsing a string to an integer, or accessing an out-of-bounds index, but *not* one that would naturally cause a nil pointer dereference). Instead of `panic`-ing, it should **return an `error`** [5].

**Instructions:**
1.  Define a function `parsePositiveInt(s string) (int, error)`.
2.  Inside `parsePositiveInt`:
    *   Attempt to convert `s` to an integer. Use `strconv.Atoi` which already returns an error.
    *   If the conversion fails, **return the error** (e.g., `return 0, err`).
    *   Additionally, if the parsed integer is not positive (i.e., less than or equal to 0), **return a new error** using `fmt.Errorf` (e.g., "input must be a positive integer, got %d") [6].
3.  In `main`, call `parsePositiveInt` with various inputs (e.g., `"123"`, `"abc"`, `"-5"`, `"0"`).
4.  **Handle the returned error** using Go's idiomatic `if err != nil` check [6]. Print appropriate messages for success or different error conditions.
5.  **In comments**, explain why returning an `error` is the **idiomatic and preferred approach** here, rather than `panic`-ing, based on Go's philosophy of error handling [2, 5].

---

### Exercise 5: Observing Developer Error Panic

**Task:**
Understand a "developer error" scenario where Go's runtime will automatically trigger a `panic`, specifically related to `nil` values.

**Instructions:**
1.  In your `main` function, declare a pointer variable (e.g., `var p *int`) and leave it uninitialized, so its **zero value is `nil`** [7].
2.  Attempt to **dereference this `nil` pointer** (e.g., `*p = 10`).
3.  **Run your program** and observe the **runtime panic**.
4.  **In a comment**, explain **why this panic occurs**, relating it to the concept of a `nil` pointer not having a valid memory address to dereference [3]. Also, based on our conversation, you can mention that attempting to call a method on a `nil` interface value would result in a similar **runtime panic** because the interface holds no concrete type or value to dispatch the method call to. This is considered an **unrecoverable "developer error"** by the Go runtime [92, our conversation].

---