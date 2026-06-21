# Step 1.3: Control Flow (Spec-Compliant Loops and Conditionals) 🔀

This step covers the mechanics of decision making and iteration in Go, emphasizing syntactic rules, block scoping, and loop variable allocation.

Official documentation:
*   [Go Spec: If statements](https://golang.org/ref/spec#If_statements)
*   [Go Spec: Switch statements](https://golang.org/ref/spec#Switch_statements)
*   [Go Spec: For statements](https://golang.org/ref/spec#For_statements)
*   [Go 1.22 Release Notes: Loop Variable Scope](https://go.dev/doc/go1.22#language)

---

## 🔍 Deep Dive 1: Conditional Statements (`if` & `else`)

The syntax of Go's `if` statement differs from other languages by omitting parentheses around the condition. However, curly braces `{}` are strictly mandatory.

### Initialization Statement Scope
Go allows an optional initialization statement to execute before evaluating the condition. Variables declared in this initialization statement are scoped **only** to the `if` block, `else if` blocks, and `else` block:
```go
if val := calculateValue(); val > limit {
    fmt.Println("Above limit:", val)
} else {
    fmt.Println("Under limit:", val) // val is still accessible here!
}
// fmt.Println(val) // ❌ Compile error: val is undefined in this scope
```
This is idiomatic for handling errors or checking values that have a narrow scope of utility.

---

## 🔍 Deep Dive 2: Switch Statements

Go switches are highly expressive and support both **expression switches** and **type switches**.

### 1. Implicit Break vs. Fallthrough
Unlike C-family languages, Go does not require a `break` statement at the end of each `case`. The compiler automatically exits the switch block once a matching case is executed. 
*   If you explicitly want fallthrough behavior (executing the next case block regardless of whether it matches), you must use the `fallthrough` keyword as the final statement in the case block.
*   **Restriction**: `fallthrough` is not allowed in type switches or as the final statement of the last case.

```go
switch day := getDay(); day {
case "Saturday", "Sunday":
    fmt.Println("Weekend")
case "Monday":
    fmt.Println("Start of work week")
    fallthrough
case "Tuesday":
    fmt.Println("Work day") // Executed if day is Monday (due to fallthrough) or Tuesday
default:
    fmt.Println("Midweek")
}
```

### 2. Expression-less Switch
If the switch expression is omitted, the compiler defaults to evaluating case conditions as boolean expressions. This is a cleaner alternative to writing long `if-else` chains:
```go
score := 85
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B") // Executed
default:
    fmt.Println("F")
}
```

---

## 🔍 Deep Dive 3: For Loops (The Unified Iteration Construct)

Go has **only one looping keyword**: `for`. However, it can represent three distinct forms of iteration:

### 1. Three-component loop (Classic)
```go
for i := 0; i < 10; i++ {
    // ...
}
```
### 2. Condition-only loop (equivalent to `while`)
```go
for condition {
    // ...
}
```
### 3. Infinite loop (equivalent to `while true`)
```go
for {
    // ...
}
```

### Labeled Break, Continue, and Goto
Go supports labels for nested control flows. You can break or continue outer loops from within inner loops:
```go
OuterLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 2 {
                break OuterLoop // Terminates the outer loop
            }
        }
    }
```

---

## 🔍 Deep Dive 4: The `for range` Clause & Variable Scope

The `for range` clause iterates over sequences: arrays, slices, strings, maps, and channels.

### 1. Value Copying Behavior
When iterating over a slice or array, the range clause returns two values: the index and a **copy** of the element value. Mutating the loop variable does not affect the source slice:
```go
items := []int{1, 2, 3}
for _, v := range items {
    v = v * 2 // ⚠️ Only modifies the local copy 'v', not the slice!
}
```

### 2. CRITICAL CHANGE: Loop Variable Scope in Go 1.22+
Prior to Go 1.22, the variables declared in the loop header (e.g. `i, v := range items`) were allocated **once** and reused across every single iteration. This caused a classic bug when capturing pointers or closures inside loops:
```go
// Pre-Go 1.22 Behavior:
var funcs []func()
for _, v := range []int{1, 2, 3} {
    funcs = append(funcs, func() {
        fmt.Println(v) // Captured reference to the single shared variable 'v'
    })
}
// Calling these functions would print: 3, 3, 3

// Go 1.22+ Behavior:
// The compiler now allocates a NEW instance of 'v' per iteration.
// Calling those same functions now prints: 1, 2, 3
```
This change eliminates the need to manually redeclare variables (e.g. `v := v`) inside the loop body.

---

## ⚠️ Common Gotchas

1.  **Iterating over Maps**: Loop iterations over maps are **non-deterministic**. The Go runtime randomizes map iteration order to prevent code from relying on stable map layouts (which can change between implementations).
2.  **String Iteration (Runes vs Bytes)**: When using `for range` on a `string`, the loop iterates over Unicode code points (`runes`), not raw bytes. It decodes UTF-8 automatically, returning the byte index and the `rune` value.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement the control flow challenges. Run:
```bash
go run .
```
Verify the output structure.
