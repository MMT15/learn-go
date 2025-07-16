--------------------------------------------------------------------------------
Flow Control
Objective: Practice using Go's flow control statements, including if/else, switch, and for loops [i, 22, 23, 24].
Tasks:
1. If/Else Statements [i, 22]:
    ◦ Declare an integer variable num and assign it a value (e.g., 10).
    ◦ Write an if/else statement that checks if num is even or odd, and print an appropriate message.
    ◦ Implement a compact if statement [i, 23] that declares a variable result within the if's scope (e.g., if result := 10 % 2; result == 0 { ... }) and prints whether the number is even or odd.
2. Switch Statements [i, 23]:
    ◦ Declare a string variable day (e.g., "Monday").
    ◦ Use a switch statement to print a message based on the day's value (e.g., "It's a weekday" for Monday-Friday, and "It's the weekend!" for Saturday-Sunday). Remember that Go's switch automatically includes a break for each case [i, 23].
    ◦ Demonstrate the fallthrough keyword [i, 24] in a switch statement. Create a scenario where one case's execution continues to the next.
    ◦ Use a switch statement without a condition (acting as switch true) [i, 24] to evaluate multiple boolean conditions and print a message.
3. For Loops [i, 24]:
    ◦ Basic for loop: Write a for loop that prints numbers from 1 to 5 [i, 25].
    ◦ for loop as a while loop: Use a for loop with only a condition (no init or post statements) [i, 26] to simulate a while loop. For example, print numbers until a counter reaches 10.
    ◦ break and continue: Write a for loop that iterates from 1 to 10. Use continue to skip printing numbers divisible by 3, and use break to stop the loop when a number greater than 7 is encountered [i, 25, 26].
    ◦ Forever loop: Implement a "forever loop" (infinite loop) using for { ... } [i, 26], and include a condition with a break statement to exit the loop after a certain number of iterations (e.g., 3).

--------------------------------------------------------------------------------