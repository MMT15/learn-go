# Exercises Chapter II: Maps

Welcome to the exercises section for the **Maps** chapter! In Go, a **map is an unordered collection of key-value pairs** [1]. It maps keys to values, where the **keys are unique within a map** while the values may not be [2]. Maps are primarily used for **fast lookups, retrieval, and deletion of data based on keys** [2], making them one of the most frequently used data structures [2].

A key characteristic of maps in Go is that they are **reference types** [3]. This means that when you assign a map to a new variable, both variables will refer to the **same underlying data structure** [3]. Consequently, **changes made through one variable will be visible to the other** [3]. The **zero value of a map is `nil`** [2]. A `nil` map has no keys, and attempting to add keys to a `nil` map will result in a runtime error [4]. Maps can be initialized using the built-in `make` function or a map literal [4].

Please read each exercise carefully and try to solve it using the knowledge gained from the chapter.

---

## Exercise 1: Map Declaration, Initialization, and Basic Operations

**Objective:** To understand how to declare, initialize, add, retrieve, and check for the existence of elements in a Go map.

### Requirements:
1.  **Declare a map** named `userScores` where keys are of type `string` (representing usernames) and values are of type `int` (representing scores) [2].
2.  **Attempt to add a key-value pair** to this initially declared (but not initialized) map (e.g., `userScores["Alice"] = 100`). **Observe the behavior** and **explain in a comment** why this happens, referencing the "zero value" property of maps [2, 4].
3.  **Initialize `userScores`** using the `make` function [4].
4.  **Add at least three key-value pairs** to `userScores` (e.g., "Alice": 100, "Bob": 85, "Charlie": 92) [5].
5.  **Initialize another map** named `productPrices` using a **map literal** with at least three `string` key (product name) to `float64` value (price) pairs (e.g., {"Laptop": 1200.50, "Mouse": 25.00, "Keyboard": 75.25}) [4].
6.  **Retrieve and print the score for "Alice"** from `userScores` [5].
7.  **Attempt to retrieve and print the score for a non-existent user** (e.g., "David") from `userScores`. **Explain in a comment** what value is returned and why [5].
8.  **Check for the existence of "Bob"** in `userScores` using the **comma-ok idiom** and print whether "Bob" exists and their score if they do [5].

### Hints:
*   Remember that a `nil` map cannot have elements added to it directly [4].
*   The `make` function is essential for initializing maps [4].
*   The **comma-ok idiom** (`value, ok := map[key]`) is crucial for checking key existence [5].

---

## Exercise 2: Updating, Deleting, Iteration, and Reference Type Behavior

**Objective:** To practice updating and deleting elements, iterating over maps, and understand their "reference type" property.

### Requirements:
1.  **Update the score for "Alice"** in `userScores` (e.g., to 105) and print the updated map [6].
2.  **Delete the entry for "Charlie"** from `userScores` and print the map again to confirm the deletion [6].
3.  **Iterate over `userScores`** using the `for...range` loop and print each key-value pair [6].
4.  **In a comment, explain** whether the order of elements printed during iteration is guaranteed or not, referencing the map's "unordered collection" property [6].
5.  **Declare a new map** named `copyOfUserScores` and **assign `userScores` to it** (e.g., `copyOfUserScores := userScores`) [3].
6.  **Modify a score in `copyOfUserScores`** (e.g., `copyOfUserScores["Bob"] = 90`) [3].
7.  **Print both `userScores` and `copyOfUserScores`** after the modification [3].
8.  **Explain in a comment** why the modification to `copyOfUserScores` *did* affect `userScores`, referencing the "reference type" property of maps as described in the sources [3].

### Hints:
*   `delete(map, key)` is the function to remove entries [6].
*   `for key, value := range myMap {}` is the standard way to iterate [6].
*   Pay close attention to the output when demonstrating the reference type behavior.

---
