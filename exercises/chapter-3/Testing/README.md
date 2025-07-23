# Go Testing Exercises

Welcome to these exercises designed to help you practice and deepen your understanding of testing in Go, drawing directly from the course material [1].

## Exercise 1: Basic Testing

**Goal**: Write a basic test for a simple mathematical function.

**Background**:
In Go, test files are declared with the **`_test` suffix** in their filename, for example, `add_test.go` for an `add.go` file [2]. Tests are typically written in a separate package, such as `math_test` for a `math` package, which helps in writing more decoupled tests [3]. The **`testing` package** from the standard library is built into Go for testing [2]. Test functions are named starting with `Test` and take an argument of type `testing.T`, which provides helpful methods for testing [3]. To run tests, you use the **`go test` command** [3].

**Task**:
1.  **Create a new package** named `calculator` (or similar) and add a file `calculator.go` to it.
2.  Inside `calculator.go`, **implement a function `Multiply(a, b int) int`** that returns the product of `a` and `b`.
3.  **Create a test file** named `calculator_test.go` in the same `calculator` package. Make sure the package declaration is `package calculator_test` [2].
4.  **Write a test function `TestMultiply(t *testing.T)`** that:
    *   Calls the `Multiply` function with some input values (e.g., `Multiply(3, 4)`).
    *   Compares the actual result with an expected value.
    *   If they do not match, use `t.Fail()` to fail the test [4].
5.  **Run your test** using the command `go test ./calculator` (assuming your calculator package is in a `calculator` directory) [3].

## Exercise 2: Table-Driven Tests

**Goal**: Refactor your basic test to use a table-driven approach for more comprehensive testing.

**Background**:
Table-driven tests are a flexible way to run multiple test cases easily. Instead of writing a separate set of assertions for each scenario, you define all your function arguments and expected variables in a slice and then iterate over that slice [5].

**Task**:
1.  **Modify your `calculator_test.go` file** from Exercise 1.
2.  **Define a struct** (e.g., `multiplyTestCase`) that holds the input arguments (`a`, `b`) and the `expected` result for the `Multiply` function [5]. Ensure this struct is declared with a **lowercase identifier** (e.g., `multiplyTestCase`) if it's not useful outside your testing logic, as lowercase identifiers are not exported and are private to the package they're defined in [6, 7].
3.  **Create a slice of `multiplyTestCase` structs** with multiple test scenarios (e.g., positive numbers, negative numbers, zero, large numbers).
4.  **Iterate over this slice** in your `TestMultiply` function. For each test case:
    *   Call the `Multiply` function with the test case inputs.
    *   Compare the actual result with the `expected` result from the test case.
    *   Report any failures using methods from `testing.T` (e.g., `t.Errorf("For %d * %d, expected %d, got %d", tc.a, tc.b, tc.expected, actual)`).
5.  **Run your tests again** with `go test ./calculator`.

## Exercise 3: Code Coverage

**Goal**: Calculate and export code coverage for your tests.

**Background**:
**Code coverage** tells you how much of your actual code is covered by your tests [6]. Go tooling provides **built-in support** for generating and viewing coverage reports [6, 8].

**Task**:
1.  **Run your tests and generate a coverage profile** for your `calculator` package using the `-coverprofile` argument [6].
    *   Command: `go test -coverprofile=coverage.out ./calculator`
2.  **View the detailed coverage report** in a more readable HTML format [8].
    *   Command: `go tool cover -html=coverage.out`
3.  **Analyze the report**: Identify any lines of code in your `Multiply` function (or any other functions you might add) that are not covered by your tests and add new test cases to cover them.

## Exercise 4: Fuzz Testing (Go 1.18+)

**Goal**: Explore fuzz testing as an automated way to find bugs by manipulating inputs.

**Background**:
**Fuzz testing**, introduced in Go 1.18, is a type of automated testing that continuously manipulates inputs to a program to find bugs. It uses coverage guidance to intelligently walk through the code being fuzzed, often reaching **edge cases that humans might miss** [8, 9].

**Task**:
1.  **Create a new function, `IsPalindrome(s string) bool`**, that checks if a string is a palindrome (reads the same forwards and backward, e.g., "madam" or "racecar"). You can put this in a new package, like `stringutil`.
2.  **Write a fuzz test function `FuzzIsPalindrome(f *testing.F)`** for your `IsPalindrome` function in `stringutil_test.go` (or `calculator_test.go` if you prefer to keep it in the same file).
    *   Add **seed corpus values** using `f.Add("racecar")`, `f.Add("hello")`, `f.Add("")` [9].
    *   Inside the `f.Fuzz` callback, call your `IsPalindrome` function and add assertions. For example, if a string `s` is a palindrome, then `IsPalindrome(s)` should be `true`. If `IsPalindrome(s)` is `true`, then `IsPalindrome(Reverse(s))` should also be `true` (you might need a `Reverse` helper function for this).
3.  **(Optional) Introduce a subtle bug** into your `IsPalindrome` function that only appears with specific inputs (e.g., mishandle strings with spaces, punctuation, or case differences, if your initial implementation doesn't account for them).
4.  **Run the fuzz test** to see if it catches your bug.
    *   Command: `go test -fuzz=. ./stringutil` (or appropriate path)
    *   You can stop fuzzing by pressing `Ctrl+C` [9].
