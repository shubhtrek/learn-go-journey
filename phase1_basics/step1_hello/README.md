# Step 1.1: Hello World & Go Tooling 🚀

Welcome to Step 1.1! Let's build a rock-solid foundation by understanding exactly what happens when you run a Go program.

---

## 🔍 Deep Dive: The Mechanics of Go Executables

Unlike Python or Node.js, Go does not run on an interpreter or a virtual machine (like Java's JVM). Go compiles directly into **native machine code** (a binary). 

When you run `go build main.go`:
1. The compiler reads your source code.
2. It parses it into an Abstract Syntax Tree (AST).
3. It optimizes the code (like removing dead code or inlining small functions).
4. It compiles it into a single executable binary specific to your OS (e.g., `.exe` on Windows, or an ELF binary on Linux).
5. The compiled binary contains the **Go Runtime** embedded inside it. This runtime handles memory management, garbage collection, and scheduling concurrent tasks.

### `go run` vs `go build`
*   `go run main.go`: Under the hood, this compiles your code to a temporary directory and immediately executes it. It's great for local development because it's fast.
*   `go build main.go`: This compiles the code and saves the binary in your current folder. You deploy this binary directly to production servers. The server doesn't even need Go installed to run it!

---

## 🔠 Code Analysis

```go
package main

import "fmt"

func main() {
    fmt.Println("👋 Namaste, Go!")
}
```

*   **`package main`**: Tells Go that this file compiles to an executable program, rather than a shared library. Every executable Go program must have a package named `main` containing a `main` function.
*   **`import "fmt"`**: The format package from Go's standard library. It handles input and output formatting.
*   **`func main()`**: The entrance point of the program.

---

## ⚠️ Common Gotchas (Avoid These!)
1.  **Missing `package main`**: If you define `package hello` instead of `package main`, `go run` will fail because it cannot find the entry point.
2.  **Unused Imports**: If you import a package and don't use it, Go will throw a compile error. This is to prevent dependency bloat.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go).
1. Implement the `PrintDetails` function to print your name and favorite snack.
2. To test it, write code inside `func main()` in `main.go` that calls `PrintDetails` with your details and run `go run main.go`.
3. Tell me when you're done, and I will manually review your implementation!

---

## 🎉 Shubham's Fun Corner 🎉

### 💍 The "Dulha" (Groom) Analogy
Why is `package main` and `func main` so special? Because it is the **Dulha** of this wedding. You can have 50 other files (relatives) in the folder, but without the Dulha (`main`), there is no wedding (executable binary)! 

### ☕ Chai Break Thought
> Java developers writing:
> `public static void main(String[] args)`
>
> Go developers writing:
> `func main()`
> 
> Work smart, not hard. Smile and give `go run .` a try!
