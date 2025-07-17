# Exercises Chapter II: Structs

Welcome to the exercises section for the **Structs** chapter! This section is designed to help you better understand the fundamental concepts of structs in Go, building on the information from the course material.

A **struct** is a user-defined type in Go that contains a collection of named fields [1]. It is essentially used to group related data together to form a single unit [1]. If you are familiar with object-oriented programming, you can think of structs as lightweight classes that support **composition** but not **inheritance** [1].

Please read each exercise carefully and try to solve it using the knowledge gained from the chapter.

---

## Exercise 1: Defining, Declaring, and Initializing a Struct

**Objective:** Understand how to define a struct and declare/initialize its instances in different ways.

### Requirements:
1.  **Define a struct** named `Person` with the following fields:
    *   `FirstName` (type `string`)
    *   `LastName` (type `string`)
    *   `Age` (type `int`)
    *   `IsStudent` (type `bool`)
2.  In the `main` function, **declare a `Person` variable** named `person1` **without explicit initialization**. Print its fields to observe their **zero values** [2].
3.  **Initialize a `Person` variable** named `person2` using a **struct literal** with all fields explicitly set (e.g., "Jane", "Doe", 30, true) [2]. Print `person2`.
4.  **Initialize a `Person` variable** named `person3` using a struct literal, but **only set `FirstName` and `Age`** (e.g., "Bob", 25) [3]. Print `person3` and observe the zero values for `LastName` and `IsStudent`.
5.  **Access and print** the `FirstName` and `Age` of `person2` [3].

### Hints:
*   Remember that Go initializes any declared variable without an explicit initial value to its **zero value** [4]. For structs, this applies to all their fields [2].
*   When initializing a struct literal, you can optionally omit field names if you provide values for all fields in the order they are declared [3]. However, using field names is generally recommended for readability [2].

---

## Exercise 2: Working with Struct Pointers

**Objective:** Demonstrate how to use pointers to structs, including their implicit dereferencing behavior.

### Requirements:
1.  **Define a struct** named `Book` with `Title` (string) and `Author` (string) fields.
2.  In the `main` function, **declare a `Book` variable** named `myBook` and initialize it with "Go Programming" and "John Doe".
3.  **Create a pointer to `myBook`** named `ptrBook` using the `&` operator [3].
4.  **Print the `Title` of `myBook`** by accessing it **through `ptrBook`**. Observe that Go automatically dereferences the pointer for field access [3].
5.  **Modify the `Author` field of `myBook`** to "Jane Smith" **using `ptrBook`** [3]. Print `myBook` to confirm the change.
6.  **Create another pointer to `Book`** named `newBookPtr` using the **`new` function** [5]. Assign "Learning Go" to its `Title` and "Alice Wonderland" to its `Author`. Print the book pointed to by `newBookPtr`.

### Hints:
*   When you have a pointer to a struct, you don't need to explicitly use the `*` operator to access its fields. Go handles this implicitly [3].
*   The `new` function allocates memory for a value of a given type and returns a pointer to it. The allocated value is initialized to its **zero value** [5].

---

## Exercise 3: Structs as Function Arguments (Pass by Value vs. Pass by Pointer)

**Objective:** Understand the implications of passing structs by value versus by pointer to functions.

### Requirements:
1.  **Define a struct** named `Product` with `Name` (string) and `Price` (float64) fields.
2.  **Create a function** `updateProductValue` that takes a `Product` (by **value**) and attempts to change its `Price`.
3.  **Create another function** `updateProductPointer` that takes a pointer to `Product` (`*Product`) and changes its `Price`.
4.  In `main`:
    *   Declare a `Product` variable `item` and initialize it with "Laptop" and 1200.0.
    *   Print `item` before any function calls.
    *   Call `updateProductValue` with `item`. Print `item` again to see if its `Price` changed.
    *   Call `updateProductPointer` with the address of `item`. Print `item` again to see if its `Price` changed this time.

### Hints:
*   Structs are **value types** in Go [6]. When you pass a struct to a function, a **new copy of the struct is created** and passed [6]. Therefore, changes made to the copy inside the function will not affect the original struct.
*   To modify the original struct inside a function, you must pass a **pointer** to the struct [5, 6].

---

## Exercise 4: Exported and Unexported Struct Fields

**Objective:** Understand Go's visibility rules for struct fields (exported vs. unexported).

### Requirements:
1.  **Define a struct** named `Configuration` with the following fields:
    *   `ServerHost` (string) - This should be **exported**.
    *   `serverPort` (int) - This should be **unexported**.
    *   `DebugMode` (bool) - This should be **exported**.
2.  In the `main` function:
    *   Create an instance of `Configuration` and initialize all its fields.
    *   **Print `ServerHost` and `DebugMode`**.
    *   **Attempt to print `serverPort`**. Observe any compilation errors or warnings.
3.  Explain in comments why `serverPort` is not accessible directly from `main` (assuming `main` is in a different package from `Configuration`'s definition, or just discuss the general rule).

### Hints:
*   In Go, any variable, function, or struct field (or even the struct itself) that starts with an **uppercase identifier** is **exported** (publicly visible) [7, 8].
*   Conversely, anything starting with a **lowercase identifier** is **unexported** (private) and only visible within the package it's defined in [7, 8].

---

## Exercise 5: Embedding vs. Composition

**Objective:** Understand the difference between embedding and composition as ways to reuse struct fields.

### Requirements:
1.  **Define a struct** named `Address` with fields `Street` (string) and `City` (string).
2.  **Define a struct** named `Customer` that **embeds** the `Address` struct [9]. Add an additional field `CustomerID` (string) to `Customer`.
3.  **Define another struct** named `Supplier` that uses **composition** for its address, meaning it will have a field `BillingAddress` of type `Address` [9]. Add an additional field `SupplierID` (string) to `Supplier`.
4.  In the `main` function:
    *   Create an instance of `Customer` and initialize its `CustomerID` and the embedded `Address` fields directly (e.g., `customer1.Street = "123 Main St"`).
    *   Create an instance of `Supplier` and initialize its `SupplierID` and its `BillingAddress` fields (e.g., `supplier1.BillingAddress.Street = "456 Oak Ave"`).
    *   Print the `Street` and `City` for both `customer1` and `supplier1`.

### Hints:
*   **Embedding** allows a struct to "inherit" the fields and methods of another struct, making them directly accessible at the top level of the embedding struct [9].
*   **Composition** involves having a field within a struct that is an instance of another struct. To access the fields of the composed struct, you must specify the field name (e.g., `supplier.BillingAddress.Street`) [9].
*   While embedding can look like inheritance, Go promotes **composition** over inheritance for code reuse [1, 9].

---