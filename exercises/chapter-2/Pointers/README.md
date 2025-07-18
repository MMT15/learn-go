# Exercises Chapter II: Pointers

Welcome to the exercises section for the **Pointers** chapter! This section is designed to help you better understand the basic concepts of pointers in Go, building on the information from the course material.

A **pointer** is a variable used to store the memory address of another variable [1]. In Go, pointers are useful for efficiently moving data without copying large amounts of information and for modifying values by reference [2].

Please read each exercise carefully and try to solve it using the knowledge gained from the chapter.

---

## Exercise 1: Declaring and Dereferencing a Basic Pointer

**Objective:** Understanding the declaration, initialization, and dereferencing of a pointer.

### Requirements:
1.  Declare an `int` variable named `num` and initialize it with the value `42`.
2.  Declare a pointer variable to `int` (`*int`) named `ptrNum` [3].
3.  Assign the memory address of the `num` variable to `ptrNum` using the **`&`** (ampersand) operator [3].
4.  Print the following:
    *   The value of the `num` variable.
    *   The memory address of the `num` variable (stored in `ptrNum`).
    *   The value `ptrNum` points to (using the **dereference operator `*`**) [4].

### Hints:
*   Remember that **`&`** is used to get the memory address of a variable [3].
*   The **`*`** operator is used to retrieve the value stored in the variable that the pointer points to, a process called **dereferencing** [4].

---

## Exercise 2: Modifying a Variable's Value through a Pointer

**Objective:** Demonstrate how to modify a variable's value through its pointer.

### Requirements:
1.  Declare a `score` variable of type `float64` and initialize it with `95.5`.
2.  Declare a pointer `ptrScore` that points to `score`.
3.  Using **ONLY** the `ptrScore` pointer (and the dereference operator `*`), **increase the value** stored at the memory address by `4.5`.
4.  Print the new value of the `score` variable to confirm the modification.

### Hints:
*   You can access and modify the value of a variable through its pointer [4].
*   Go does not support pointer arithmetic as in C/C++ [5]. Value modification is done through dereferencing and re-assignment.

---

## Exercise 3: Using Pointers as Function Arguments

**Objective:** Understanding the concept of "pass by reference" using pointers as function arguments to modify a value from outside the function.

### Requirements:
1.  Create a function named `incrementValue` that takes an argument of type **pointer to `int`** (`*int`) [4].
2.  Inside the `incrementValue` function, increment the value that the pointer points to by `1`.
3.  In the `main` function:
    *   Declare an `int` variable `counter` and initialize it with `10`.
    *   Print the value of `counter` before calling the function.
    *   Call the `incrementValue` function, **passing the memory address** of the `counter` variable.
    *   Print the value of `counter` after the function call to see if it has been modified.

### Hints:
*   You can use pointers as arguments for a function when you need to **pass data by reference** [4]. This means that changes made within the function will affect the original variable [6].
*   Go is "pass by value", but by passing the address (a pointer), you can simulate "pass by reference" to allow the function to modify the original data.

---

## Exercise 4: Zero Value for Pointers and the `new` Function

**Objective:** Exploring the zero value of a pointer and using the `new` function to initialize pointers.

### Requirements:
1.  Declare a pointer variable to `string` (`*string`) named `uninitializedPtr` **without explicitly initializing it**.
2.  Print the value of `uninitializedPtr`. Observe that the zero value for pointers is **`nil`** [3].
3.  Declare a pointer variable to `bool` named `newBoolPtr` using the **`new` function** [5].
4.  Assign the value `true` to the memory location indicated by `newBoolPtr`.
5.  Print the value that `newBoolPtr` points to.

### Hints:
*   **`nil`** is the predeclared identifier in Go that represents the zero value for pointers, interfaces, channels, maps, and slices [3].
*   The **`new` function** takes a type as an argument, allocates enough memory to accommodate a value of that type, and returns a pointer to it. The initial value that `new` points to is the zero value of that type [5].

---

## Exercise 5: Comparing Pointers

**Objective:** Understanding how pointers can be compared in Go.

### Requirements:
1.  Declare two `int` variables `a` and `b` and initialize them with `10` and `20` respectively.
2.  Declare two pointers `ptrA` and `ptrB` to point to `a` and `b` respectively.
3.  Compare `ptrA` and `ptrB` for equality and print the result.
4.  Make `ptrB` now point to `a`.
5.  Compare `ptrA` and `ptrB` for equality again and print the new result.

### Hints:
*   You can compare two pointers of the same type for equality using the `==` operator [2].
*   Two pointers are equal if both are `nil` or if they point to the same memory address [2].

---