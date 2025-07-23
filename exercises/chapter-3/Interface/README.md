# Exercises Chapter III: Interfaces

Welcome to the exercises section for the **Interfaces** chapter! In Go, an **interface is an abstract type** that is defined by a **set of method signatures** [1]. It defines the **behavior** for similar types of objects [1, 2]. A crucial aspect of Go interfaces is that they are **implemented implicitly**; a type automatically satisfies an interface when it has **all the methods of the interface** [3]. There is no explicit `implements` keyword like in other languages [3].

The convention for naming interfaces is to use the **"-er" suffix** (e.g., `Reader`, `Writer`, `PowerDrawer`) [2]. Interfaces are incredibly powerful as they help us **decouple our types**, meaning the implementer doesn't need to know anything about the concrete type, only that it satisfies the required behavior [3]. The **zero value of an interface is `nil`** [4], and an interface value can internally be thought of as a **tuple consisting of a value and a concrete type** [5].

Please read each exercise carefully and try to solve it using the knowledge gained from the chapter.

---

## Exercise 1: Implicit Interface Implementation and Polymorphism

**Objective:** To understand how types implicitly satisfy interfaces and how functions can operate on different concrete types via an interface.

### Requirements:
1.  **Define an interface** named `Measurable` with a single method signature: `Area() float64`.
2.  **Define a struct** named `Rectangle` with fields `Width` and `Height` (both `float64`).
3.  **Implement the `Area()` method for `Rectangle`** so that it calculates and returns the area (`Width * Height`).
4.  **Define another struct** named `Circle` with a field `Radius` (`float64`).
5.  **Implement the `Area()` method for `Circle`** so that it calculates and returns the area (`math.Pi * Radius * Radius`).
6.  **Create a function** named `PrintMeasurement` that takes an argument of type `Measurable` and prints the area using the `Area()` method.
7.  In your `main` function:
    *   **Create instances of `Rectangle` and `Circle`**.
    *   **Call `PrintMeasurement`** with both the `Rectangle` and `Circle` instances.
8.  **Explain in a comment** how `Rectangle` and `Circle` can be passed to `PrintMeasurement` even though they are different types, referencing the **implicit implementation** and **behavior** aspects of interfaces [1, 3].

### Hints:
*   You might need to import the `math` package for `math.Pi`.
*   Remember that Go doesn't need an `implements` keyword.

---

## Exercise 2: Empty Interfaces, Type Assertion, and Type Switch

**Objective:** To practice using the **empty interface (`interface{}`)** to handle values of unknown types and to extract concrete values using **type assertion** and **type switch**.

### Requirements:
1.  **Declare a slice of empty interfaces** named `items` (e.g., `[]interface{}`).
2.  **Populate `items`** with values of various types, such as an `int`, a `string`, a `bool`, a `float64`, and an instance of `Rectangle` or `Circle` from Exercise 1.
3.  **Iterate over the `items` slice** using a `for...range` loop.
4.  Inside the loop, use a **type switch** statement to check the concrete type of each `item`:
    *   If it's an `int`, print: "Found an integer: [value]".
    *   If it's a `string`, print: "Found a string: [value]".
    *   If it's a `bool`, print: "Found a boolean: [value]".
    *   If it's a `float64`, print: "Found a float: [value]".
    *   If it's a `Rectangle` (or `Circle`), print: "Found a shape (Rectangle/Circle) with area: [area]". You'll need to use a **type assertion** to call its `Area()` method [6].
    *   For any other type, print: "Found an unknown type: [value]".
5.  **Demonstrate a safe type assertion** (using the comma-ok idiom) outside the loop for one specific element in `items` (e.g., try to assert an `int` to a `string` and print the `ok` value and zero value if it fails) [6].
6.  **Explain in a comment** why empty interfaces are useful, particularly in scenarios where data types are heterogeneous or unknown, referencing the `fmt.Println` example mentioned in the sources [7].

### Hints:
*   The syntax for type switch is `switch v := i.(type) {}` [4].
*   The comma-ok idiom for type assertion is `value, ok := interfaceVar.(ConcreteType)` [6].

---

## Exercise 3: Interface Embedding and `nil` Interface Values

**Objective:** To explore **interface embedding** and understand the **zero value of an interface** and its implications.

### Requirements:
1.  **Define a base interface** named `Notifier` with a method `Notify(message string)`.
2.  **Define a second interface** named `Alerter` that **embeds `Notifier`** and adds an additional method `RaiseAlert(level int, message string)`.
3.  **Create a struct** named `SystemMonitor` with no fields.
4.  **Implement the `Notify()` method for `SystemMonitor`** that prints a notification message.
5.  **Implement the `RaiseAlert()` method for `SystemMonitor`** that prints an alert message with a given level.
    *   By implementing both, `SystemMonitor` implicitly satisfies both `Notifier` and `Alerter` due to embedding [4].
6.  In your `main` function:
    *   **Create an instance of `SystemMonitor`**.
    *   **Declare a variable of type `Notifier`** and assign the `SystemMonitor` instance to it. Call its `Notify()` method.
    *   **Declare a variable of type `Alerter`** and assign the `SystemMonitor` instance to it. Call both its `Notify()` and `RaiseAlert()` methods.
7.  **Demonstrate the zero value of an interface**:
    *   Declare an interface variable (e.g., `var myInterface Alerter`) without initializing it.
    *   Print this `nil` interface variable.
    *   **Attempt to call a method** (e.g., `myInterface.Notify("test")`) on this `nil` interface and **explain in a comment** the runtime behavior you observe, relating it to the concept of a `nil` interface value [4].

### Hints:
*   Interface embedding works similar to struct embedding [4].
*   Calling a method on a `nil` interface value where the concrete type is also `nil` will result in a panic.

---