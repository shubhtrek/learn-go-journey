# Step 1.3: Control Flow (The One Loop) 🔀

Control flow dictates how code branches and loops. Go keeps it minimalist.

---

## 🔍 Deep Dive: If, Switch, and the Unified Loop

### 1. `if / else`
Go's `if` statement does not require parentheses `()` around the condition, but curly braces `{}` are strictly mandatory.
Go also supports an **initialization statement** inside the `if`:
```go
if val := getVal(); val > 10 {
    // val is only accessible inside this if/else block scope!
}
```
This is a very common pattern in Go for scope limitation.

### 2. `switch`
Go switches do not require a `break` statement at the end of each case. It only executes the matching case. If you want to fall through to the next case, you must write `fallthrough` explicitly.

### 3. `for` (The Only Loop)
Go has no `while` or `do-while` loops. The `for` loop does it all:
*   **Classic loop**: `for i := 0; i < 10; i++ {}`
*   **While loop equivalent**: `for condition {}`
*   **Infinite loop**: `for {}`

---

## ⚠️ Common Gotchas
1.  **Brace Placement**: Go compiler automatically inserts semicolons. Therefore, the opening brace `{` must be on the same line as the `if` or `for` statement.
    ```go
    if x > 5
    { // ❌ COMPILE ERROR: syntax error: unexpected newline
    }
    ```

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go).
1. Implement the `FizzBuzz` function.
2. In `main.go`, write a loop that calls your `FizzBuzz` function from 1 to 20 and prints the outputs.
3. Message me for manual review when done!

---

## 🎉 Shubham's Fun Corner 🎉

### 🍛 The "Indian Wedding Host" Analogy
Go having only one loop (`for`) is like a wedding host standing at the exit door saying, *"Khana kha ke hi jaana!"* (Eat before you leave!) on repeat until you actually eat. 
There is no `while` loop, no `do-while` loop. Just one `for` loop that can dress up to look like anything you want.

### ☕ Chai Break Thought
> While loop? Do-while loop? 
> Go compiler: *We don't do that here.*
Keep it simple. Take a deep breath, stretch, and conquer the FizzBuzz challenge!
