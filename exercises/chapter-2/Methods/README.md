# Exercises Chapter II: Methods

Welcome to the exercises section for the **Methods** chapter! In Go, a method is a function with a special **receiver** argument [1]. While Go is not strictly an object-oriented programming language and doesn't have classes or inheritance, you can define methods on types [1]. The receiver argument appears between the `func` keyword and the method name [1], allowing you to access the instance of the type using the receiver variable, similar to `this` in other languages [2].

A crucial concept when working with methods is the distinction between **value receivers** and **pointer receivers**:
*   **Value receivers:** When a method has a value receiver, it operates on a **copy of the value** passed to it. Therefore, any modifications made to the receiver inside the method **will not be visible to the caller** [3].
*   **Pointer receivers:** Methods with pointer receivers **can modify the value** to which the receiver points. Such modifications **are visible to the caller** of the method [3]. Go also provides syntactic sugar, implicitly dereferențiind pointerul pentru accesul la câmpuri și apelurile de metodă [4].

Methods are not limited to structs; they can also be used with non-struct types [4]. They can also help avoid naming conflicts, as the same method name can exist for multiple receivers tied to different types [5].

Please read each exercise carefully and try to solve it using the knowledge gained from the chapter.

---

## Exercise 1: Defining a Method with a Value Receiver

**Objective:** Understand the basic syntax for defining a method and observe the behavior of methods with **value receivers**, specifically that they operate on a copy of the original value.

### Requirements:
1.  **Define a struct** named `Rectangle` with the following fields:
    *   `Length` (type `float64`)
    *   `Width` (type `float64`)
2.  **Define a method** named `Area()` on the `Rectangle` struct. This method should take a **value receiver** (e.g., `r Rectangle`). It should calculate and return the area of the rectangle (`Length * Width`).
3.  **Define another method** named `Scale(factor float64)` on the `Rectangle` struct. This method should also take a **value receiver**. Inside this method, attempt to update the `Length` and `Width` of the receiver by multiplying them with the `factor`.
4.  In the `main` function:
    *   Create an instance of `Rectangle` named `myRect` and initialize its `Length` to 10.0 and `Width` to 5.0.
    *   **Print the initial `myRect`** instance.
    *   Call `myRect.Area()` and **print the calculated area**.
    *   Call `myRect.Scale(2.0)` to attempt to double its dimensions.
    *   **Print `myRect` again** after calling `Scale()`. Observe and comment on why the dimensions of `myRect` did not change, relating it to the concept of value receivers [3].

### Hints:
*   Remember that a method's receiver type determines whether the method operates on a copy or the original value [3].
*   For a value receiver, the syntax is `func (r Rectangle) MethodName(...)`.

---

## Exercise 2: Methods with Pointer Receivers for Modification

**Objective:** Demonstrate how to modify the original struct using methods with **pointer receivers**.

### Requirements:
1.  Using the `Rectangle` struct defined in Exercise 1.
2.  **Define a new method** named `SetDimensions(newLength, newWidth float64)` on the `Rectangle` struct. This method **must use a pointer receiver** (e.g., `r *Rectangle`). Inside this method, update the `Length` and `Width` fields of the receiver to `newLength` and `newWidth` respectively.
3.  In the `main` function:
    *   Create another instance of `Rectangle` named `anotherRect` and initialize it with arbitrary values (e.g., Length 7.0, Width 3.0).
    *   **Print the initial `anotherRect`** instance.
    *   Call `anotherRect.SetDimensions(20.0, 10.0)`.
    *   **Print `anotherRect` again** after calling `SetDimensions()`. Confirm that its dimensions have now been successfully modified.
    *   Explain the key difference in behavior between this method and the `Scale` method from Exercise 1, referencing the properties of pointer receivers [3].

### Hints:
*   For a pointer receiver, the syntax is `func (r *Rectangle) MethodName(...)`.
*   Go automatically handles dereferencing when accessing fields of a struct through a pointer receiver, so you can use `r.Length` directly instead of `(*r).Length` [4].

---

## Exercise 3: Methods on Non-Struct Types

**Objective:** Understand that methods are not exclusive to structs and can be defined on any user-defined type.

### Requirements:
1.  **Define a new type** named `Celsius` that is an alias for `float64` (e.g., `type Celsius float64`).
2.  **Define a method** named `ToFahrenheit()` on the `Celsius` type. This method should take a **value receiver** (`c Celsius`) and convert the temperature from Celsius to Fahrenheit using the formula: `F = C * 9/5 + 32`. Return the result as a `float64`.
3.  **Define another method** named `IncreaseBy(delta float64)` on the `Celsius` type. This method should take a **pointer receiver** (`c *Celsius`) and increase the temperature by `delta`.
4.  In the `main` function:
    *   Declare a `Celsius` variable named `tempC` and initialize it to 25.0.
    *   Call `tempC.ToFahrenheit()` and **print the Fahrenheit equivalent**.
    *   **Print the initial `tempC` value**.
    *   Call `tempC.IncreaseBy(5.0)` to increase its value.
    *   **Print `tempC` again** after calling `IncreaseBy()`. Observe and confirm that the `tempC` variable itself was modified.

### Hints:
*   Even though `Celsius` is based on `float64`, it's a distinct type that can have its own methods.
*   The same rules for value vs. pointer receivers apply to non-struct types.

---