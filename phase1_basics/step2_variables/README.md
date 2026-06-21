# Step 1.2: Variables, Types, & Zero Values 📦

Let's dive into how Go stores, reads, and manages data in memory.

---

## 🔍 Deep Dive: Go's Static Type System & Memory

Go is a **statically typed** and **strongly typed** language:
*   **Statically typed**: The compiler must know the type of every variable at compile time. You cannot change a variable's type after declaration.
*   **Strongly typed**: Go will not automatically convert types for you (no implicit type casting). You cannot add an `int` and a `float64` without converting one of them first.

### Variable Declaration Syntaxes
1.  **Verbose (`var`)**:
    ```go
    var count int = 10
    ```
    Useful when you want to define a variable without giving it a value immediately, or when defining package-level variables.
2.  **Short Assignment (`:=`)**:
    ```go
    count := 10
    ```
    This is the idiomatic way inside functions. Go infers the type automatically based on the value on the right.

### The Zero Value Concept
In Go, declaring a variable without an initial value automatically assigns it its **Zero Value**:
*   `int`, `float64`: `0`
*   `string`: `""` (empty string)
*   `bool`: `false`
*   Pointers, slices, maps, channels: `nil`

This design prevents bugs caused by "garbage values" in memory (a common issue in C/C++).

---

## ⚠️ Common Gotchas
1.  **Implicit Casting Failure**:
    ```go
    var x int = 5
    var y float64 = 4.5
    // z := x + y // ❌ COMPILE ERROR: invalid operation (mismatched types int and float64)
    z := float64(x) + y // ✅ Correct: explicit cast
    ```
2.  **Short Variable Redeclaration**: You cannot use `:=` to redeclare a variable in the same scope unless at least one new variable is being introduced on the left side.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go).
1. Implement `GetVariables` to return an integer `42`, float `3.14`, string `"Go is awesome"`, and boolean `true`.
2. Update your `main.go` to call this function and print the returned values.
3. Let me know when you've written the code, and I'll review it!

---

## 🎉 Shubham's Fun Corner 🎉

### 📋 The "Strict Indian Parents" Analogy
Go's static typing is like typical strict parents. If you declared that you want to be an **Engineer** (`int`), you *cannot* suddenly decide to become a **Chef** (`float64`) midway without a major family meeting (explicit casting)! Go will literally stop compiling and tell you: "Beta, focus on one type."

### ☕ Chai Break Thought
In Go, variables are never left with random junk in memory. They get a supportive zero-value. If only our plans had this kind of automatic backup system!


### 👑 Marathi Swag: No Timepass with Types!
*   Go is super strict. It’s like having a **khadus (strict) uncle** who won't let you mix Chai and Coffee. If you declare an `int`, it stays `int`.
*   If you try to add an `int` and a `float` directly, Go will scream: **"Aata bagha, mismatched types!"** You have to convert them explicitly: `float64(your_int)`.
*   Zero Values: Go handles them automatically. Even if you forget to initialize, Go sets it to `0` or `""`. No empty-pocket stress!
