--------------------------------------------------------------------------------
Functions
Objective: Practice defining, calling, and utilizing various features of functions in Go, including parameters, return values, anonymous functions, closures, variadic functions, the init function, and the defer keyword [i, 27, 28, 29, 30, 31, 32].
Tasks:
1. Simple Function Declaration and Call [i, 27]:
    ◦ Define a simple function named greet that takes no arguments and prints "Hello, Go functions!" to the console.
    ◦ Call this greet function from your main function.
2. Functions with Parameters [i, 27]:
    ◦ Create a function named add that takes two integer parameters, a and b. Inside the function, print the sum of a and b.
    ◦ Call add with example values.
    ◦ Demonstrate shorthand parameter declaration by creating another function, subtract, that takes two int parameters and prints their difference.
3. Functions with Return Values [i, 27]:
    ◦ Modify the add function to return the sum of a and b instead of printing it. Print the result returned by add in main.
    ◦ Implement multiple return values [i, 27]: Create a function calc that takes two integers and returns both their sum and their product. Call calc and print both results.
4. Named Returns [i, 28]:
    ◦ Re-implement the calc function (from Task 3) using named return values [i, 28]. Observe how you can use a naked return [i, 28].
5. Functions as Values and Anonymous Functions [i, 28, 29]:
    ◦ Declare a variable op and assign your add function (the one that returns the sum) to it [i, 28]. Call op with two numbers and print the result.
    ◦ Create and immediately execute an anonymous function that prints "This is an anonymous function." [i, 29].
6. Closures [i, 29]:
    ◦ Write a function makeCounter that returns another function (a closure) [i, 29]. The returned function should increment an internal counter variable (from makeCounter's scope) each time it's called.
    ◦ Demonstrate by creating a counter instance and calling it multiple times.
7. Variadic Functions [i, 30]:
    ◦ Create a function sumAll that takes an arbitrary number of integers (...int) [i, 30] and returns their total sum.
    ◦ Call sumAll with different numbers of arguments (e.g., sumAll(1, 2, 3) and sumAll(10, 20, 30, 40)).
    ◦ Hint: Remember that fmt.Println is itself a variadic function! [i, 30]
8. The init Function [i, 30, 31, 32]:
    ◦ Define an init function in your main.go file that prints "Init function executed." [i, 31].
    ◦ Observe that the init function is executed before your main function [i, 31].
    ◦ (Optional): Create a second init function in the same file to observe the order of execution [i, 31].
9. The defer Keyword [i, 32]:
    ◦ Create a function cleanup that prints "Cleanup finished!"
    ◦ In your main function, call defer cleanup() before any other prints. Observe when "Cleanup finished!" is printed [i, 32].
    ◦ Demonstrate the "last-in, first-out" (LIFO) behavior of defer statements [i, 32] by deferring multiple functions that print different messages.

--------------------------------------------------------------------------------