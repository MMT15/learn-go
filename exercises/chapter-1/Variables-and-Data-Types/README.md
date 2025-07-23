---

# **Exercises: Chapter I - Variables and Data Types**

This section provides a set of exercises to help you solidify your understanding of variables, constants, and various data types in Go. For each exercise, create a new `.go` file (e.g., `exercise1.go`) and run it using the `go run` command.

---

## **Exercise 1: Variable Declarations and Initialization**

**Objective:** Practice different ways to declare and initialize variables in Go.

**Tasks:**

1.  Declare an integer variable named `age` without initializing it and then print its **zero value**.
2.  Declare and initialize a string variable named `name` with your name in one line, and then print its value.
3.  Declare two integer variables, `x` and `y`, with values `10` and `20` respectively, using a **shorthand declaration** within the `main` function. Print their sum.
    *   **Hint**: Remember that shorthand declaration `:=` only works inside function bodies.
4.  Declare a boolean variable named `isGoFun` and initialize it to `true`. Print its value.
5.  Declare multiple variables `a`, `b`, and `c` of type `string` on a single line and assign them values "Go", "is", and "awesome" respectively. Print them concatenated.

## **Exercise 2: Constants**

**Objective:** Understand how to declare and use constants.

**Tasks:**

1.  Declare a constant named `PI` with the value `3.14159`.
2.  Declare a constant named `Greeting` with the value "Hello, Go!".
3.  Try to reassign a new value to `PI` after its declaration and observe the compilation error.
4.  Print both constants.

## **Exercise 3: Numeric Data Types and Operators**

**Objective:** Explore different numeric types and apply various operators.

**Tasks:**

1.  Declare an `int` variable `num1` with value `15` and a `float64` variable `num2` with value `4.5`.
2.  Perform and print the result of the following **arithmetic operations** between `num1` and `num2` after **type conversion**:
    *   Addition (`+`)
    *   Subtraction (`-`)
    *   Multiplication (`*`)
    *   Division (`/`)
    *   **Note**: You will need to explicitly convert `num1` to `float64` or `num2` to `int` before performing operations. The source recommends using `int` for integers unless there's a specific reason.
3.  Declare a `uint8` variable named `byteValue` and assign it the value `65` (which is the ASCII value for 'A'). Print its value.
4.  Declare an `int32` variable named `runeValue` and assign it the Unicode code point for the smiley face character `😊`. Print its value.
    *   **Hint**: You might need to use `fmt.Printf` with a specific **annotation verb** to see the character correctly, which will be covered in the next chapter. For now, just print the integer value.
5.  Declare two integer variables `p` and `q` with values `7` and `3` respectively. Perform and print the result of the following **bitwise operations**:
    *   Bitwise AND (`&`)
    *   Bitwise OR (`|`)
    *   Left shift (`<<`) by 1.
6.  Use the **increment (`++`) and decrement (`--`) operators** on an integer variable `counter` initialized to `5` and print the value after each operation.

## **Exercise 4: Boolean Logic and Comparison**

**Objective:** Understand boolean data types and their operators.

**Tasks:**

1.  Declare two boolean variables: `isRaining` set to `true` and `isCold` set to `false`.
2.  Print the result of the following **logical operations**:
    *   `isRaining && isCold` (AND)
    *   `isRaining || isCold` (OR)
    *   `!isRaining` (NOT)
3.  Declare two integer variables `val1` with value `100` and `val2` with value `200`.
4.  Print the result of the following **comparison operations**:
    *   `val1 == val2` (Equality)
    *   `val1 != val2` (Inequality)
    *   `val1 < val2` (Less than)
    *   `val1 >= val2` (Greater than or equal to)

## **Exercise 5: Type Conversion and Aliases vs. Defined Types**

**Objective:** Differentiate between type conversion, alias types, and defined types.

**Tasks:**

1.  Declare an `int` variable `integerVal` with value `42`.
2.  Convert `integerVal` to a `float32` and store it in `floatVal`. Print both `integerVal` and `floatVal` with their **types**.
    *   **Hint**: You can use `fmt.Printf` with the `%T` **annotation verb** to print the type.
3.  Define an **alias type** `MyString` for `string`. Declare a variable `aliasVar` of type `MyString` and assign it a string literal. Print its value and type.
4.  Define a **defined type** `UserID` with an **underlying type** of `int`. Declare a variable `myID` of type `UserID` and assign it an integer value.
5.  Try to assign an `int` variable directly to `myID` and observe the compilation error, illustrating that defined types are not interchangeably used with their underlying types.
6.  Convert an `int` variable to `UserID` explicitly and print the result.

---