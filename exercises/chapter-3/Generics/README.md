# Go Generics Exercises

Welcome to these exercises designed to help you practice and deepen your understanding of **Generics** in Go, a much-awaited feature released with **Go version 1.18** [1, 2].

**Background on Generics**:
Generics, also known as **parameterized types**, allow programmers to **write code where the type can be specified later** because the type isn't immediately relevant [2]. This is particularly useful in Go because **method overriding is not allowed**, often leading to the need for creating separate functions for different types (e.g., `sumInt`, `sumFloat64`, `sumString`) [2]. Generics can **drastically reduce code duplication** in such cases [3, 4].

A generic function is defined with a **type parameter** (e.g., `T`) and a **constraint**, which is an **interface that allows any type implementing the interface** [5]. Since Go 1.18, you can use `any` as a constraint, which is equivalent to the empty interface (`interface{}`) [5]. Go 1.18 also introduced **type inference**, allowing you to call generic functions without explicitly passing type arguments, making your code less verbose [6].

## Exercise 1: Basic Generic Function

**Goal**: Implement a basic generic function that works with multiple numeric types without needing separate functions for each.

**Background**:
As learned, generics help avoid code duplication when the logic is similar across different types [2, 3]. For simple scenarios, the `any` constraint can be used, although it has limitations [5, 6].

**Task**:
1.  **Create a new package** (e.g., `utils`) and a file `utils.go`.
2.  **Implement a generic function `PrintSlice[T any](slice []T)`** that takes a slice of any type `T` and prints each element on a new line.
3.  In your `main` package, **create slices of different types** (e.g., `[]int`, `[]string`, `[]float64`).
4.  **Call your `PrintSlice` function** with these different slices, demonstrating its reusability. Notice how Go's type inference works [6].
5.  **Run your program** to observe the output.

## Exercise 2: Generic Function with Custom Constraints

**Goal**: Implement a generic function that performs an arithmetic operation, requiring a custom type constraint.

**Background**:
While `any` is versatile, it **does not support operators** like `+` [6]. To perform operations specific to certain types (e.g., addition), you need to define a **custom interface** that acts as a constraint, specifying the **type set** of allowed types (e.g., `int | float64 | string`) [3, 6].

**Task**:
1.  **Modify your `utils` package** from Exercise 1.
2.  **Define a custom interface `Numeric`** that includes `int` and `float64` in its type set. For example: `type Numeric interface { int | float64 }` [3].
3.  **Implement a generic function `Sum[T Numeric](numbers ...T) T`** that takes a variadic number of arguments of type `T` (constrained by `Numeric`) and returns their sum.
4.  In your `main` package, **call your `Sum` function** with both `int` and `float64` values.
5.  **Run your program** and verify the sums.
6.  **(Optional challenge)** Try to add `string` to the `Numeric` interface and see if you can concatenate strings using the `+` operator within the `Sum` function. What happens? (Hint: The `+` operator behaves differently for strings and numbers).

## Exercise 3: Using Standard Library Constraints

**Goal**: Utilize the `constraints` package for common type constraints.

**Background**:
Go's standard library provides the `constraints` package, which defines a set of **useful constraints to be used with type parameters** [3]. For example, the `Ordered` constraint is helpful for types that can be ordered (numbers, strings). The `~` token is also introduced in Go 1.18, meaning "the set of all types whose **underlying type** is string" (e.g., `~string`) [3].

**Task**:
1.  **Install the `constraints` package** if you haven't already. You might need to use `go install golang.org/x/exp/constraints@latest`.
2.  **Modify your `utils` package** (or create a new one).
3.  **Implement a generic function `Max[T constraints.Ordered](a, b T) T`** that takes two arguments of an `Ordered` type `T` and returns the larger of the two.
4.  In your `main` package, **call your `Max` function** with different ordered types (e.g., `int`, `float64`, `string`).
5.  **Run your program** and confirm it correctly identifies the maximum value.
6.  **(Optional challenge)** Research and use another constraint from the `constraints` package (e.g., `constraints.Integer`, `constraints.Float`) for a new generic function.

## Further Exploration: When to Use Generics

Generics are a powerful addition to Go, but they should be **used sparingly** [4]. It's often advised to **start simple and only write generic code once you have written very similar code at least 2 or 3 times** [4].

Some common use cases where generics shine include [4]:
*   Functions that operate on **arrays, slices, maps, and channels**.
*   **General-purpose data structures** like stacks or linked lists.
*   To **reduce code duplication**.