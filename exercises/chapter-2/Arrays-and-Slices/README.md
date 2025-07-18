# Exercises Chapter II: Arrays and Slices

Welcome to the exercises section for the **Arrays and Slices** chapter! In Go, **arrays** are fixed-size collections of elements of the same type [1]. Their elements are stored sequentially and can be accessed using their index [1, 2]. A key property of arrays in Go is that their **length is part of their type**, meaning arrays of different lengths are distinct types [3]. Moreover, arrays are **value types**; when an array is assigned to a new variable or passed to a function, the entire array is copied [3, 4]. This implies that modifications to the copied array will not affect the original [4].

While arrays are useful, their fixed size can be inflexible [4]. This is where **slices** come in. A slice is a segment of an array, providing more power, flexibility, and convenience [4]. A slice consists of three components: a pointer reference to an underlying array, its length (the number of elements it contains), and its capacity (the maximum size the segment can grow to) [5]. Unlike arrays, slices are **reference types** [6]. This means that modifying the elements of a slice will modify the corresponding elements in the referenced underlying array [6]. Slices can be created using the `make` function or by slicing an existing array or another slice [7].

Please read each exercise carefully and try to solve it using the knowledge gained from the chapter.

---

## Exercise 1: Understanding Arrays (Fixed Size and Value Type)

**Objective:** To understand the declaration, initialization, and the "value type" property of Go arrays.

### Requirements:
1.  **Declare an array** named `grades` of type `[8]int` (an array of 5 integers).
2.  **Initialize `grades`** with arbitrary integer values using an array literal (e.g., `[8]int{90, 85, 92, 78, 95}`).
3.  **Print the `grades` array** to observe its contents.
4.  **Declare another array** named `studentGrades` and assign `grades` to it (e.g., `studentGrades := grades`).
5.  **Modify the first element** of `studentGrades` (e.g., `studentGrades = 100`).
6.  **Print both `grades` and `studentGrades`** after the modification.
7.  **Explain in a comment** why the modification to `studentGrades` did not affect `grades`, referencing the "value type" property of arrays as described in the sources [3, 4].

### Hints:
*   Remember that the length of an array is part of its type [3].
*   When you assign one array to another, a full copy is made [4].

---

## Exercise 2: Working with Slices (Dynamic Nature and Reference Type)

**Objective:** To understand slice declaration, initialization, the concepts of length and capacity, and their "reference type" behavior.

### Requirements:
1.  **Declare a slice** named `numbers` of type `[]int` (a slice of integers).
2.  **Initialize `numbers`** using the `make` function to have a length of 3 and a capacity of 5 (e.g., `make([]int, 3, 5)`).
3.  **Print `numbers`, its length (`len(numbers)`), and its capacity (`cap(numbers)`)**. [5]
4.  **Initialize the elements** of `numbers` (e.g., `numbers = 10`, `numbers[9] = 20`, `numbers[10] = 30`).
5.  **Declare another slice** named `moreNumbers` and assign `numbers` to it (e.g., `moreNumbers := numbers`).
6.  **Modify the second element** of `moreNumbers` (e.g., `moreNumbers[9] = 25`).
7.  **Print both `numbers` and `moreNumbers`** after the modification.
8.  **Explain in a comment** why the modification to `moreNumbers` *did* affect `numbers`, referencing the "reference type" property of slices as described in the sources [6].

### Hints:
*   The `make` function is used to create slices, maps, and channels [7].
*   `len()` gives the number of elements in the slice, and `cap()` gives the capacity of the underlying array [5].

---

## Exercise 3: Expanding Slices with `append` and `copy`

**Objective:** To practice using the built-in `append` and `copy` functions for slices, understanding how they interact with the underlying array.

### Requirements:
1.  **Create a slice** named `data` using a slice literal: `[]string{"apple", "banana", "cherry"}`.
2.  **Append two new elements** to `data` (e.g., "date", "elderberry") using the `append` function. Store the result in a new variable, `newData`, as `append` returns a new slice [11].
3.  **Print `data` and `newData`**, along with their lengths and capacities. Observe if `data` changed (it shouldn't, as `append` returns a new slice).
4.  **Create another slice** named `sourceSlice` with values `{"one", "two", "three", "four"}`.
5.  **Create an empty destination slice** named `destinationSlice` with a length of 2 and a capacity of 4 using `make` (e.g., `make([]string, 2, 4)`).
6.  **Copy elements from `sourceSlice` to `destinationSlice`** using the `copy` function. Print the number of elements copied [12].
7.  **Print `sourceSlice` and `destinationSlice`** after the copy operation.
8.  **Explain in a comment** what happened during the `copy` operation regarding the number of elements copied, referring to how `copy` works with the length of the destination slice [12].

### Hints:
*   The `append` function can reallocate the underlying array if the capacity is insufficient, returning a new slice [6].
*   The `copy` function copies elements up to the length of the *smaller* of the two slices [12].

---