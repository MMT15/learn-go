# Go `sync` Package Exercises

Welcome to the `sync` package! When you're dealing with **goroutines** [1], which run in the same address space, **access to shared memory must be synchronized** [2, 3]. The `sync` package provides **useful primitives** to help you manage this concurrency safely and efficiently [3]. This set of exercises will guide you through using some of the most common and essential tools in this package.

## Getting Started

For each exercise, create a new `.go` file and implement your solution. Experiment with different scenarios and observe the results!

## Exercises

### Exercise 1: `sync.WaitGroup` - Waiting for Goroutines

**Objective:** Understand how to use `sync.WaitGroup` to ensure your main goroutine waits for a collection of other goroutines to finish [3]. This is crucial for controlling the lifecycle of concurrent tasks.

**Task:**
1.  In your `main` function, declare a `sync.WaitGroup` variable.
2.  Define a function (e.g., `worker`) that takes an integer `id` and a `*sync.WaitGroup` as parameters. This function should print a message like "Worker [id] starting..." and "Worker [id] finished." after a short `time.Sleep` (e.g., 500ms).
3.  Inside `main`, launch 5 `worker` goroutines:
    *   Before launching each goroutine, call **`wg.Add(1)`** to increment the counter of goroutines the `WaitGroup` needs to wait for [4].
    *   Inside the `worker` function, make sure to call **`defer wg.Done()`** at the beginning. This decrements the `WaitGroup` counter when the goroutine finishes [4].
4.  After launching all worker goroutines, call **`wg.Wait()`** in `main`. This will block the `main` goroutine until all workers have called `Done()` [4].
5.  **Observe**: What happens if you remove `wg.Wait()`? What if you forget `wg.Add(1)` or `wg.Done()`? (Hint: The program might exit early or deadlock). Remember that a `WaitGroup` **must not be copied after first use**, and if passed to functions, it should be done **by a pointer** [5].

### Exercise 2: `sync.Mutex` - Preventing Data Races

**Objective:** Learn to use `sync.Mutex` to protect shared resources and prevent **data races** [6]. A `Mutex` is a mutual exclusion lock that prevents multiple processes from entering a **critical section** of data concurrently [5, 7].

**Task:**
1.  Define a `Counter` struct with an integer field `value` and a `sync.Mutex` field (e.g., `mu sync.Mutex`).
2.  Implement an `Increment()` method for the `Counter` struct. This method should:
    *   Call **`c.mu.Lock()`** before incrementing the `value` [8].
    *   Call **`defer c.mu.Unlock()`** after incrementing, to ensure the lock is released even if a panic occurs [7].
3.  In your `main` function:
    *   Create a `Counter` instance.
    *   Use a `sync.WaitGroup` (from Exercise 1) to manage a set of goroutines.
    *   Launch 1000 goroutines, and each goroutine should call `counter.Increment()` 100 times.
4.  After all goroutines complete, print the final `counter.value`.
5.  **Observe**: What is the final value? Try running the program *without* the `Mutex` (comment out `Lock()` and `Unlock()`) and see the different, inconsistent results due to the **data race** [8]. Remember, a `Mutex` **must not be copied after first use** [8].

### Exercise 3: `sync.RWMutex` - Reader/Writer Lock

**Objective:** Understand `sync.RWMutex`, which allows multiple readers to hold a lock concurrently, but only one writer. It's preferable for data that is mostly read [9].

**Task:**
1.  Modify your `Counter` struct from Exercise 2 to use `sync.RWMutex` instead of `sync.Mutex` (e.g., `rwmu sync.RWMutex`).
2.  Update the `Increment()` method to use **`c.rwmu.Lock()`** and **`defer c.rwmu.Unlock()`** for writing [9].
3.  Add a `GetValue()` method to the `Counter` struct. This method should:
    *   Call **`c.rwmu.RLock()`** before reading the `value` [10].
    *   Call **`defer c.rwmu.RUnlock()`** after reading [10].
    *   Return the current `value`.
4.  In your `main` function:
    *   Launch a few goroutines that *write* to the counter using `Increment()`.
    *   Launch many more goroutines that *read* the counter using `GetValue()`. Add `time.Sleep` calls to simulate work and allow concurrency to be visible.
5.  **Observe**: How does `RWMutex` behave differently from a plain `Mutex` when multiple readers are active? When does a read operation block?

### Exercise 4: `sync.Once` - Ensuring Single Execution

**Objective:** Use `sync.Once` to ensure that a particular function or block of code is executed **only once**, even if called concurrently by multiple goroutines [11].

**Task:**
1.  Declare a `sync.Once` variable.
2.  Define a function, `performSetup()`, that prints "Performing one-time setup..." and then increments a global counter (or a counter within a struct that holds the `sync.Once`).
3.  In your `main` function:
    *   Use a `sync.WaitGroup` to launch 10 goroutines.
    *   Each goroutine should call **`once.Do(performSetup)`** [11].
4.  Print the final value of the global counter after all goroutines complete.
5.  **Observe**: Despite `performSetup` being called 10 times concurrently through `once.Do`, how many times was the actual setup message printed and the counter incremented? It should be exactly once [11].

### Exercise 5: `sync.Map` - Concurrency-Safe Map

**Objective:** Understand how `sync.Map` provides a concurrency-safe map suitable for specific use cases, avoiding the need for manual locking with a standard Go map [12].

**Task:**
1.  In your `main` function, declare a `sync.Map` variable.
2.  Use a `sync.WaitGroup` to launch 5 goroutines for writing and 5 goroutines for reading concurrently.
3.  **Writing Goroutines (e.g., `storeWorker`):**
    *   Each goroutine should generate a unique key-value pair (e.g., "keyN", N) and store it in the `sync.Map` using **`m.Store(key, value)`** [13].
    *   Add a small `time.Sleep` to simulate work.
4.  **Reading Goroutines (e.g., `loadWorker`):**
    *   Each goroutine should attempt to load a specific key (or a range of keys) from the `sync.Map` using **`m.Load(key)`** [13].
    *   Print the loaded value.
    *   Use **`m.LoadOrStore(key, value)`** to demonstrate conditionally storing a value if it doesn't exist [13].
5.  After all goroutines complete, iterate over the `sync.Map` using **`m.Range(func(key, value any) bool)`** to print all stored key-value pairs [13].
6.  **Observe**: Does the map behave as expected under concurrent writes and reads without explicit `Mutex` calls? Recall that `sync.Map` is optimized for cases where entries are written once but read many times, or when goroutines operate on disjoint sets of keys [14]. A `sync.Map` **must not be copied after first use** [14].

---

## Hints and Tips

*   The Go proverb **"Don't communicate by sharing memory, share memory by communicating"** [1] is paramount. While the `sync` package helps when you *must* share memory, channels are often a more idiomatic Go solution.
*   Always consider **error handling** in real-world concurrent applications, as panics can occur if synchronization primitives are misused (e.g., unlocking an unlocked mutex).
*   When using `WaitGroup`, ensure that `Add` is called *before* the goroutine starts to guarantee correct accounting [4].
*   `defer` is your friend for ensuring locks are released (e.g., `defer mu.Unlock()`) and `WaitGroup.Done()` is called [15].
*   An empty `select {}` will block forever [3]. This is not directly `sync` package, but relevant to concurrency control.
*   You might need `time.Sleep` (temporarily) in `main` to see output from goroutines, or (preferably) `sync.WaitGroup` to properly wait for them [2].