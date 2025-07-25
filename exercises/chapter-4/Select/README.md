# Go `select` Statement Exercises

Alright, let's dive into the **`select` statement** in Go! This is a really powerful construct that lets you work with **multiple channel operations simultaneously** [1]. Think of it as a `switch` statement, but for channels – it **blocks and waits** until one of its cases can run, then executes that specific case [1]. If multiple channels are ready at the same time, `select` **picks one at random** [1].

These exercises will help you get comfortable with `select` and its behavior.

## Getting Started

For each exercise, create a new `.go` file and implement your solution. Don't be afraid to experiment with different timings to see how `select` reacts!

## Exercises

### Exercise 1: Basic Multi-Channel Selection

**Objective:** Understand how `select` waits for and picks the first ready channel.

**Task:**
1.  In your `main` function, create two **unbuffered channels** of type `string`, let's call them `ch1` and `ch2`.
2.  Launch **two goroutines**:
    *   One goroutine should send "Message from channel 1" to `ch1` after a `time.Sleep` of 1 second.
    *   The other goroutine should send "Message from channel 2" to `ch2` after a `time.Sleep` of 500 milliseconds (0.5 seconds).
3.  In your `main` goroutine, use a **`select` statement** with two `case` branches, one for receiving from `ch1` and one for `ch2`.
4.  Print the message you receive.
5.  **Observe**: Which message do you receive? Try changing the sleep durations to see how it affects the outcome.

### Exercise 2: `select` with a `default` Case

**Objective:** Learn how the `default` case in `select` prevents blocking and runs immediately if no other channel operation is ready.

**Task:**
1.  In your `main` function, create an **unbuffered channel** of type `int`.
2.  Use a **`select` statement** that tries to receive from this channel.
3.  Add a **`default` case** to the `select` statement that prints "No channel operation ready."
4.  **Run the program**. What happens?
5.  Now, try launching a goroutine that sends a value to the channel *after* a short delay (e.g., 100 milliseconds), but keep the `select` statement to execute immediately. What output do you get now? (This highlights the non-blocking nature of `default`).

### Exercise 3: Iterating with `select` and Channel Closure

**Objective:** Understand how `select` works within a loop and handles channel closure.

**Task:**
1.  Create a `main` function.
2.  Declare an **unbuffered channel** of type `int`, say `dataCh`.
3.  Launch a **goroutine** that sends integers from 1 to 5 to `dataCh`, with a small `time.Sleep` between each send (e.g., 200ms).
4.  After sending all 5 integers, **close `dataCh`** in this goroutine.
5.  In the `main` goroutine, use a **`for` loop** that contains a **`select` statement**.
6.  Inside the `select`, have a `case` to receive from `dataCh`. When you receive a value, print it.
7.  Also, have a `case` that checks if `dataCh` is closed (by checking the second `ok` return value: `val, ok := <-dataCh`). If it's closed and `ok` is `false`, print "Channel closed. Exiting." and use `return` or `break` to exit the loop/function.
8.  **Observe**: How does `select` manage to receive all values and then gracefully exit after the channel is closed?

### Exercise 4: Implementing Timeouts with `select`

**Objective:** Use `select` to implement a timeout for a channel operation. This is a common pattern for handling long-running or potentially stuck operations.

**Task:**
1.  In your `main` function, create an **unbuffered channel** of type `string`, say `resultCh`.
2.  Launch a **goroutine** that simulates a long-running operation: it should send "Operation complete!" to `resultCh` after a `time.Sleep` of 3 seconds.
3.  In your `main` goroutine, use a **`select` statement** to either:
    *   Receive the message from `resultCh`.
    *   Receive a signal from `time.After(2 * time.Second)` (This channel sends a value after the specified duration. *Note: `time.After` is a function from Go's standard `time` package, not explicitly detailed in the `select` source [1], but it's a very common and idiomatic way to implement timeouts with `select` statements, conceptually similar to how `context` handles deadlines and cancellations [2, 3].*)
4.  Print an appropriate message depending on whether the operation completed or timed out.
5.  **Experiment**: Change the `time.After` duration to be longer than the operation's sleep and observe the difference.

---

## Hints and Tips

*   Remember the core rule: a `select` statement **blocks** until one of its cases is ready to proceed [1].
*   An **empty `select {}` will block forever** [4].
*   When using `select` in a loop, always consider how you will **exit the loop** once your work is done or channels are closed to prevent infinite blocking or panics.
*   You might need `time.Sleep` in your `main` goroutine to prevent it from exiting before other goroutines have a chance to send/receive, especially with unbuffered channels, or use `sync.WaitGroup` as you've seen in the `Channels` and `Sync Package` lessons [4, 5].