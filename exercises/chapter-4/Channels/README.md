# Go Channels Exercises

This set of exercises is designed to help you practice and deepen your understanding of **channels** in Go, a fundamental concept for **concurrent programming** [1-3]. Channels serve as **communication pipes between goroutines**, allowing them to **share memory by communicating**, adhering to the important Go proverb [3, 4].

## Getting Started

For each exercise, create a new `.go` file and implement the solution. Feel free to experiment!

## Exercises

### Exercise 1: Basic Unbuffered Channel Communication

**Objective:** Understand how to send and receive data using an unbuffered channel between two goroutines [5].

**Task:**
1.  Create a `main` function.
2.  Inside `main`, create an unbuffered channel of type `string` [5].
3.  Launch a new **goroutine** that sends the message "Hello from goroutine!" to the channel [5, 6].
4.  In the `main` goroutine, **receive the message from the channel** and print it to the console [5].
5.  **Observe the synchronous nature of unbuffered channels**: sending will block until a receiver is ready, and vice versa [7, 8].

### Exercise 2: Buffered Channel Behavior

**Objective:** Explore the behavior of **buffered channels** and their capacity [7].

**Task:**
1.  Create a `main` function.
2.  Declare a **buffered channel** of type `int` with a capacity of 2 [7].
3.  **Send two integer values** (e.g., 10 and 20) into the channel [7]. Note that this won't block immediately because the buffer has space.
4.  Try to **send a third value**. Observe if the program blocks or panics (it should block until a value is received, as sends to a full buffered channel block) [7].
5.  Receive and print all values from the channel [5].

### Exercise 3: Directional Channels in Functions

**Objective:** Learn how to enforce send-only or receive-only behavior for channels when passed as **function parameters** [9].

**Task:**
1.  Create a `main` function.
2.  Declare an unbuffered channel of type `int`.
3.  Write a function `producer(ch chan<- int)` that takes a **send-only channel** as an argument [9]. Inside, send a value (e.g., 5) to the channel.
4.  Write another function `consumer(ch <-chan int)` that takes a **receive-only channel** as an argument [9]. Inside, receive and print a value from the channel.
5.  In `main`, launch `producer` and `consumer` in separate goroutines, passing the channel appropriately.
6.  Use `time.Sleep` (temporarily) or `sync.WaitGroup` to ensure `main` waits for the goroutines to complete, addressing the non-deterministic scheduling behavior observed previously [6, 10].

### Exercise 4: Closing Channels and Range Iteration

**Objective:** Understand how to **close a channel** and **iterate over it using `for...range`** [9, 11].

**Task:**
1.  Create a `main` function.
2.  Declare an unbuffered channel of type `string`.
3.  Launch a goroutine that sends three different string messages to the channel (e.g., "Alpha", "Beta", "Gamma") [5].
4.  After sending all messages, **close the channel** inside this goroutine [9].
5.  In the `main` goroutine, use a **`for...range` loop** to receive and print messages from the channel [11].
6.  Observe how the `for...range` loop automatically exits when the channel is closed and all values have been received [11].

### Exercise 5: Channel with `sync.WaitGroup`

**Objective:** Combine channels with `sync.WaitGroup` to **properly synchronize goroutines**, ensuring all concurrent operations complete before the main function exits [6, 10]. This reinforces the pattern of handling goroutine completion.

**Task:**
1.  Create a `main` function.
2.  Initialize a `sync.WaitGroup` [10].
3.  Create an unbuffered channel of type `int`.
4.  Define a constant `numWorkers = 3`.
5.  Launch `numWorkers` goroutines:
    *   Each goroutine should call `wg.Add(1)` (or `wg.Add(numWorkers)` once before the loop).
    *   Each goroutine should then call `defer wg.Done()` to signal completion [12].
    *   Each goroutine should receive an integer from the channel, double it, and print the result.
6.  In the `main` goroutine, send `numWorkers` integer values to the channel (one for each worker).
7.  After sending all values, **close the channel** [9].
8.  Finally, call `wg.Wait()` in `main` to ensure all worker goroutines have finished before the program exits [12].

---

## Hints and Tips

*   Remember that channels are **reference types** and should be passed by value, not pointer, to functions unless you specifically want to modify the channel itself (which is rare) [not explicitly in sources, but common Go practice derived from slice/map behavior].
*   The `go` keyword is used to start a new **goroutine** [6].
*   When a channel is closed, receiving from it will immediately return the **zero value** of its type if no more values are present, and the `ok` boolean will be `false` [13].
*   `fmt.Println()` is a **variadic function** and can take multiple arguments [14].
*   Consider using `go run .` from your project directory to run your main file [15].