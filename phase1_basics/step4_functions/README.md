# Step 1.4: Functions, Signatures & Memory Evaluation 📞

This step covers the mechanics of function declarations, parameter passing strategy, multiple return values, variadic operations, and lexical closures.

Official documentation:
*   [Go Spec: Function Declarations](https://golang.org/ref/spec#Function_declarations)
*   [Effective Go: Named Results](https://go.dev/doc/effective_go#named-results)
*   [Go Spec: Passing Arguments](https://golang.org/ref/spec#Calls)

---

## 🔍 Deep Dive 1: Function Signatures & Parameter Lists

A function declaration binds an identifier to a function type.
```go
func Calculate(x int, y int) (int, error) {
    // ...
}
```
*   **Parameter shorthand**: If consecutive parameters share the same type, you can omit the type for the preceding parameters:
    ```go
    func Add(a, b, c int) int { // a, b, and c are all integers
        return a + b + c
    }
    ```
*   **Signature Equality**: Two functions have the same type (signature) if they have the same parameter types and the same result types.

---

## 🔍 Deep Dive 2: Return Parameters & Naked Returns

Go supports returning multiple values, which is the idiomatic mechanism for returning results alongside errors.

### Named Return Values
The return parameters of a function can be named. If named, they are treated as local variables declared at the top of the function and initialized to their zero values:
```go
func Split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // "Naked" return
}
```
### Naked Return Rules & Effective Go Guidelines
*   A `return` statement without arguments (a **naked return**) returns the current values of the named return variables.
*   **Effective Go Recommendation**: While named return values are highly readable when documenting what the function returns (acting as documentation in package APIs), naked returns should be restricted to **short, simple functions**. In longer functions, naked returns degrade readability because the reader must scan the function body to trace the current state of the return variables.

---

## 🔍 Deep Dive 3: Variadic Parameters (`...T`)

A function can accept a variable number of arguments by designating the final parameter as variadic:
```go
func SumAll(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```
*   Under the hood, the compiler converts the variadic parameter `nums` into a slice of type `[]int` inside the function body.
*   **Passing a slice**: If you already have a slice, you cannot pass it directly to a variadic parameter. You must unpack it using the `...` operator:
    ```go
    prices := []int{10, 20, 30}
    result := SumAll(prices...) // Unpacks the slice
    ```

---

## 🔍 Deep Dive 4: First-Class Functions & Closures

Functions in Go are first-class citizens. You can assign functions to variables, pass them as arguments to other functions, and return them.

### Closures
A closure is an anonymous function that references variables from outside its lexical scope. The function "binds" to these variables, preserving their state even after the outer function has returned:
```go
func CounterGenerator() func() int {
    count := 0
    return func() int {
        count++ // Mutates and preserves 'count' from the outer scope
        return count
    }
}

func main() {
    next := CounterGenerator()
    fmt.Println(next()) // 1
    fmt.Println(next()) // 2
}
```

---

## 🔍 Deep Dive 5: Pass-by-Value Semantics (Strictly enforced)

A critical rule of Go: **Everything is passed by value.** 

When you pass an argument to a function, Go always creates a **copy** of the value. There is no such thing as pass-by-reference at the runtime level.

### 1. Passing Primitive Types
Passing an `int` or `string` copies the data. Modifying it inside the function has no effect on the original variable.

### 2. Passing Pointers
When you pass a pointer (`*int`), the pointer address itself is **copied**. However, because the copy points to the same memory address as the original pointer, dereferencing the copied pointer lets you mutate the original memory.

### 3. Slices, Maps, and Channels
Beginners often mistakenly believe slices and maps are passed by reference because modifying them inside a function affects the original data. 
*   **The Reality**: A slice variable is a struct containing a pointer to a backing array, a length, and a capacity. When passed, this header struct is **copied by value**. However, because the copied header contains the *same pointer* to the backing array, modifications to elements inside the slice affect the shared backing array. 
*   However, if you append elements inside the function and cause the slice to reallocate, the changes to the header (length and capacity) are **not** visible to the caller because the header copy was mutated, not the original.

---

## ⚠️ Common Gotchas

1.  **Naked Return Shadowing**: If you declare a local variable inside a function with the same name as a named return variable, it shadows it. A naked return will still return the original named return variables (which will remain at their zero values or previous values), ignoring the local variable's state.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement functions that demonstrate named returns, closures, and parameter passing. Run:
```bash
go run .
```
Verify compilation.
