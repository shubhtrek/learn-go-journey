# Step 1.3: Control Flow (Only ONE Loop!) 🔀

This step details Go's conditional execution constructs (`if`, `switch`) and iteration mechanisms (`for`). 

We will learn this using systems engineering precision, but with a side of **Chai-Samosa** to keep it delicious!

Official documentation:
*   [Go Spec: If statements](https://golang.org/ref/spec#If_statements)
*   [Go Spec: Switch statements](https://golang.org/ref/spec#Switch_statements)
*   [Go Spec: For statements](https://golang.org/ref/spec#For_statements)

---

## ☕ Chai-Samosa FizzBuzz Analogy

Other languages have multiple loop structures like `while`, `do-while`, `forEach`, and `for`. Go developers decided that was too much clutter. In Go, we **only have one loop: `for`**. We can configure it to behave like any other loop!

To test your loop logic, we use the classic **Chai-Samosa** (FizzBuzz) challenge:
*   Loop from 1 to 50:
    *   If the number is divisible by 3, print **"Chai"** (instead of Fizz).
    *   If divisible by 5, print **"Samosa"** (instead of Buzz).
    *   If divisible by both, print **"Chai-Samosa"** (instead of FizzBuzz).
    *   Otherwise, just print the number.

---

## 🔍 Deep Dive 1: Conditional Scoping (`if`/`else`)

Go `if` statements can include a short initialization statement executed before the conditional check. Variables declared in this initialization block are only in scope until the end of the `if`/`else` blocks:

```go
if val := getStatus(); val == "OK" {
    fmt.Println("Success:", val) // val is visible here
} else {
    fmt.Println("Failed:", val)  // val is visible here too
}
// fmt.Println(val) // ❌ Compile error: val is undefined here (out of scope!)
```
*Hinglish tip*: Custom scope management is built-in. Variable scope block ke bahaar nahi jayega, meaning memory is cleaned up quickly!

---

## 🔍 Deep Dive 2: Switch Rules & `fallthrough`

Go's `switch` is more flexible than in C/Java:
*   **Implicit Break**: Go automatically breaks out of a switch block after executing a matching case. You do not need to write `break` at the end of every case.
*   **Multiple Values**: A single case can evaluate multiple comma-separated values:
    ```go
    switch day {
    case "Saturday", "Sunday":
        fmt.Println("Weekend!")
    }
    ```
*   **The `fallthrough` Keyword**: If you explicitly want to fall through to the next case (ignoring that case's conditional check), use the `fallthrough` statement:
    ```go
    switch x {
    case 1:
        fmt.Println("One")
        fallthrough
    case 2:
        fmt.Println("Two") // Executed if x == 1 because of fallthrough!
    }
    ```
*   **Expressionless Switch**: You can write a switch without an expression, which acts as a clean chain of `if-else` blocks:
    ```go
    switch {
    case x < 0:
        fmt.Println("Negative")
    case x > 0:
        fmt.Println("Positive")
    }
    ```

---

## 🔍 Deep Dive 3: The Unified `for` Loop & Go 1.22+ Memory Update

Go uses a single keyword `for` to implement all loops:

```go
// 1. Classic Three-Component Loop
for i := 0; i < 10; i++ {}

// 2. While Loop equivalent (Conditional check only)
for x < 100 {}

// 3. Infinite Loop (for ever!)
for {}
```

### The `for range` Value Copying gotcha
When iterating over slices or arrays using `range`, Go returns two values: index and a **copy** of the element's value. Modifying the loop variable does not update the parent slice:
```go
arr := []int{1, 2, 3}
for _, val := range arr {
    val = val * 10 // Only updates the temporary local variable 'val'!
}
fmt.Println(arr) // Prints [1, 2, 3]
```

### ⚠️ Go 1.22+ Loop Variable Allocation Update
Prior to Go 1.22, variables declared in a `for` loop were allocated **once** and reused across all iterations. This caused major concurrency bugs when sharing loop variable references with goroutines:
```go
// Pre-Go 1.22 behavior:
for _, val := range arr {
    go func() {
        fmt.Println(val) // ⚠️ All goroutines printed the final value because they shared the same pointer!
    }()
}
```
**In Go 1.22+**, the runtime allocates a **new instance of the loop variable on every iteration**, preventing this class of race conditions automatically!

---

## 👑 Marathi Swag: Loop cha Ekach Pattern!
*   Other languages have 4 types of loops, but Go says: **"Khup jatra nako!"** (No crowded mess). Only `for` loop is enough!
*   `fallthrough` is there if you want to skip brakes, but be careful.
*   Go 1.22 has updated loop variables. No more copy-paste closure bugs. **Ekdum kadak solution!**
*   Open [practice.go](./practice.go) to write the Chai-Samosa challenge. Run `go run .` to test it!
