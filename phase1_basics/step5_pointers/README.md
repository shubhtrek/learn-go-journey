# Step 1.5: Pointers Demystified 📍

Pointers let you share and mutate memory locations directly. Let's make pointers simple.

---

## 🔍 Deep Dive: Stack vs Heap & Pointer Operations

Every variable is stored in a memory location with a specific address. A **pointer** is a variable that stores the memory address of another variable.

### Operators:
*   `&` (Address-of): Gets the memory address of a variable.
*   `*` (Dereference / Value-at): Accesses or modifies the value stored at the address a pointer is pointing to.

```go
var x int = 10
var p *int = &x // p stores the address of x
fmt.Println(*p) // Prints 10 (dereferences p)
*p = 20         // Changes x to 20 in memory
```

### Stack vs Heap (Escape Analysis)
In languages like C, you must manually allocate memory on the Heap (`malloc`) and free it. 
In Go, the compiler performs **Escape Analysis**:
*   If a variable's lifetime can be determined to remain within the function scope, it is allocated on the **Stack** (super fast).
*   If the variable outlives the function (e.g., returned as a pointer), the compiler "escapes" it to the **Heap**, where it will eventually be cleaned up by the Garbage Collector.

---

## ⚠️ Common Gotchas
1.  **Dereferencing a `nil` pointer**: If a pointer doesn't point to any memory address (it is `nil`), trying to dereference it (`*p`) will crash your program at runtime (Panic). Always check if a pointer is `nil` before dereferencing if there's a chance it might be empty!

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go).
1. Implement `DoubleValue` to double the value at the given pointer.
2. Implement `SwapValues` to swap the values of two integers using pointers.
3. Test your functions inside `main.go`.
4. Message me for manual review!
