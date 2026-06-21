# Step 1.4: Functions & Memory Evaluation 📞

This step covers Go function signatures, named and naked return variables, lexical closures, and the strict pass-by-value execution model.

We will learn this with systems-level precision, but with the classic **Xerox Copy vs. Cupboard Key** analogy!

Official documentation:
*   [Go Spec: Function Declarations](https://golang.org/ref/spec#Function_declarations)
*   [Go Spec: Closures](https://golang.org/ref/spec#Function_literals)
*   [Effective Go: Named Result Parameters](https://go.dev/doc/effective_go#named-results)

---

## ☕ The Xerox Copy vs. Cupboard Key Analogy

Go is strictly **pass-by-value**. Every time you pass a variable to a function, Go copies the value in memory.

### The Analogy:
*   **Pass-by-value (Xerox Copy)**: You have a premium sheet of notes (a variable). A function wants to read it. Instead of giving them the original notes, you make a **Xerox copy** and hand it over. If they write doodles or change the numbers on their copy, your original notes remain clean! This is Go's default.
*   **Pointers (Cupboard Key)**: What if you want the function to actually mutate the original value? You don't pass the whole cupboard; you pass a duplicate **Key** to the cupboard (the pointer containing the memory address). Using that key, the function opens your cupboard and modifies the contents inside.

---

## 🔍 Deep Dive 1: Function Signatures & Named Returns

Go functions support multiple return values, named returns, and variadic parameters.

### Named and Naked Returns
You can name the return variables in the function signature. They act as local variables initialized to their zero values. If you use named returns, a simple `return` statement (known as a **naked return**) will return their current values:

```go
func SplitSalary(total int) (basic int, allowance int) {
    basic = (total * 70) / 100
    allowance = total - basic
    return // Returns basic and allowance automatically!
}
```

#### Code Quality Rule (Naked Returns)
While naked returns reduce boilerplates, they can hurt readability in large functions because the reader has to look back at the function signature to see what is being returned. 
**Best Practice**: Only use naked returns in small, simple functions (less than 10 lines).

---

## 🔍 Deep Dive 2: Lexical Closures (Function Literals)

Go supports anonymous functions, which can form **closures**:
*   A closure is a function value that references variables from outside its body.
*   The function can access and modify the referenced variables; they survive even after the enclosing scope has exited.

```go
func SequenceGenerator() func() int {
    i := 0 // Local variable
    return func() int {
        i++ // Accesses 'i' from lexical outer scope
        return i
    }
}
```
*Hinglish tip*: Function value capture. Generator call karne pe local state preserve ho rahi hai, memory heap pe store ho jayegi automatic!

---

## 🔍 Deep Dive 3: Proof of Pass-by-Value

A common point of confusion is whether slices and maps are passed by reference. They are not.
*   **The Rule**: Everything in Go is passed by value.
*   **The Reality**: When you pass a slice or map to a function, Go makes a copy of the **header struct** (which contains a pointer to the backing array). 
    *   Since the copy of the header still points to the *same* backing array, modifying elements (e.g. `slice[0] = 99`) affects the caller's slice.
    *   However, if you reassign the slice parameter itself (e.g. `slice = append(slice, 10)`), it changes the copied header's pointer and length, leaving the caller's slice header untouched!

---

## 👑 Marathi Swag: Return cha Dhadaka!
*   In Go, functions can return multiple values at the same time. Like ordering a **Thali**—you get Samosa, Roti, and Gulab Jamun together!
*   Naked returns: Clean but don't abuse them in big functions, or your team will get confused.
*   Remember the Xerox analogy: Go always copies values by default. If you want changes to stick, you must use pointers (the cupboard key). **Khara shahanpan pointers vaparnyat aahe!**
*   Open [practice.go](./practice.go) to solve the challenge.
