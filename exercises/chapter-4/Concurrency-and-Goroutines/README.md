# Chapter IV Exercises: Concurrency & Goroutines

This section provides exercises to solidify your understanding of **Concurrency** and **Goroutines** in Go, drawing upon the concepts introduced in Chapter IV of the course [1]. These exercises will help you grasp the fundamental principles and practical application of concurrent programming in Go.

## Introduction to Concurrency

### Exercise 1: Concurrency vs. Parallelism

**Task:** Explain in your own words the fundamental difference between **Concurrency** and **Parallelism** [2, 3].

*   Provide a real-world example for each concept to illustrate your explanation [2, 3].
*   **Hint:** Consider Rob Pike's insightful quote: "Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once" [3]. Reflect on how this quote helps differentiate the two concepts.

### Exercise 2: Understanding CSP

**Task:** Briefly describe the **Communicating Sequential Processes (CSP)** model [3], which Go relies on for its concurrency [3, 4].

*   How does this model guide the design of concurrent Go programs, especially regarding communication between processes, without explicitly mentioning channels yet [4]?

## Working with Goroutines

### Exercise 3: Your First Goroutine

**Task:**
1.  Create a simple Go function, for example, `func greet(message string)`. Make this function print the `message` to the console [5].
2.  From your `main` function, call `greet` as a **goroutine** by using the `go` keyword before the function call [6].
3.  Run the program. What do you observe about the output [6]? Specifically, why might you not see the `greet` message printed completely or at all [6]?
    *   **Insight:** Remember that the `main` function itself runs as a goroutine, and if the main goroutine exits, the program terminates without waiting for other goroutines to complete [6, 7].
4.  Modify your `main` function to include a `time.Sleep` call (e.g., `time.Sleep(time.Second)`) after launching the goroutine. Re-run the program. What changes do you observe in the output now [6, 8]?
    *   **Note:** While `time.Sleep` can make a goroutine's output visible, it is **not an ideal or reliable solution** for synchronizing goroutines in real applications, as it's a fixed delay and doesn't guarantee completion [8]. More robust synchronization mechanisms (like `sync.WaitGroup` or `channels`) will be covered in later sections of Chapter IV [1, 8, 9].

### Exercise 4: Multiple Goroutines and Non-Determinism

**Task:**
1.  Create a function `worker(id int)` that takes an integer `id` as an argument. Inside this function, print a message indicating "Worker [id] starting...", then simulate some work by sleeping for a random short duration (e.g., `time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)` - remember to import `math/rand` and `time`), and finally print "Worker [id] finished."
2.  In your `main` function, launch **5** instances of the `worker` function, each as a separate goroutine with a unique `id` [5, 7].
3.  As in the previous exercise, use `time.Sleep` in `main` to ensure all goroutines have a chance to complete before the `main` goroutine exits [6, 8].
4.  Run the program multiple times. What do you observe about the order of the "starting" and "finished" messages [7]? Why is the order typically not consistent across runs [7, 8]?

### Exercise 5: Identifying a Race Condition (Conceptual)

**Task:**
1.  Consider the following Go pseudocode snippet, where `counter` is a global variable:
    ```go
    var counter int = 0

    func increment() {
        // Assume this function is called by multiple goroutines concurrently
        temp := counter // Read the current value of counter
        temp = temp + 1 // Perform an operation
        counter = temp  // Write the new value back to counter
    }
    ```
2.  If multiple goroutines call the `increment` function concurrently, what type of **concurrency issue** might arise [8, 10]? Use the specific term mentioned in the course materials [10].
3.  Explain in detail why the final value of `counter` might not be the expected sum of all increments (e.g., if 10 goroutines call `increment`, why might `counter` not be 10) [8, 10]?
4.  Relate this problem back to the important Go proverb shared in the course: **"Don't communicate by sharing memory, share memory by communicating."** [7]. How does this proverb suggest a different approach to solving such problems in Go [4, 7]?

---