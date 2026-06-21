# Step 1.2: Variables, Types, Constants & Zero Values 📦

This step focuses on how Go manages data in memory, its strong type system, constants evaluation, and compile-time type safety.

Official documentation:
*   [Go Spec: Variable Declarations](https://golang.org/ref/spec#Variable_declarations)
*   [Go Spec: Short Variable Declarations](https://golang.org/ref/spec#Short_variable_declarations)
*   [Go Spec: Types](https://golang.org/ref/spec#Types)
*   [Go Spec: Constants](https://golang.org/ref/spec#Constants)
*   [Go Spec: The Zero Value](https://golang.org/ref/spec#The_zero_value)

---

## 🔍 Deep Dive 1: Variable Declarations & Scope

Go offers multiple declaration syntaxes depending on location and scope requirements.

### 1. The `var` Statement
The `var` statement is the standard form of declaring variables. It specifies the variable name and type.
```go
var x int
var name string = "Shubham"
```
*   If a type is provided, the compiler allocates memory initialized to the type's **zero value** (unless an initial value is assigned).
*   If an initial value is assigned without an explicit type, the compiler infers the type based on the value on the right-hand side.
*   `var` declarations can exist at both the **package-level** (global scope) and **block-level** (inside functions).

### 2. Short Variable Declaration (`:=`)
Inside functions (block scope), the short variable declaration syntax can be used:
```go
count := 10 // Infers count is of type int
```
*   **Restriction**: `:=` is *not* allowed at the package level.
*   **Redeclaration Rule**: You can use `:=` to redeclare variables in the same block only if:
    1. The redeclaration occurs in the same block.
    2. The variable being redeclared is of the same type.
    3. At least one *new* variable is introduced on the left side of the `:=` operator.
    ```go
    x := 1
    // x := 2 // ❌ Compile error: no new variables on left side
    x, y := 2, 3 // ✅ Allowed: y is new
    ```

---

## 🔍 Deep Dive 2: Go's Static Type System

Go is **strongly and statically typed**. Unlike languages like JavaScript or Python, variable types are fixed at compile-time. Unlike C or C++, Go does not perform **implicit type coercion**.

### Basic Types
| Category | Types | Memory Size | Details |
| :--- | :--- | :--- | :--- |
| **Integers (Signed)** | `int8`, `int16`, `int32`, `int64`, `int` | 1, 2, 4, 8 bytes; `int` is 32 or 64 bits | `int` is the default integer type. |
| **Integers (Unsigned)** | `uint8` (byte), `uint16`, `uint32`, `uint64`, `uint` | 1, 2, 4, 8 bytes; `uint` is 32 or 64 bits | `byte` is an alias for `uint8`. |
| **Unicode Code Points** | `rune` | 4 bytes | `rune` is an alias for `int32`. |
| **Floating-Point** | `float32`, `float64` | 4, 8 bytes | `float64` is the default floating-point type. |
| **Complex Numbers** | `complex64`, `complex128` | 8, 16 bytes | Real and imaginary parts. |
| **Boolean** | `bool` | 1 byte | `true` or `false`. |
| **Strings** | `string` | 16 bytes (on 64-bit OS) | Immutable sequence of bytes containing UTF-8 characters. |

### Strict Conversions
You cannot mix types in assignments or expressions, even if they share the same underlying architecture size:
```go
var a int32 = 100
var b int64 = 200
// c := a + b // ❌ Compile error: mismatched types int32 and int64

c := int64(a) + b // ✅ Correct: explicit type conversion
```

---

## 🔍 Deep Dive 3: Constants & The `iota` Generator

Constants in Go are declared with the `const` keyword. They are values known at compile-time and cannot be altered.

### Untyped Constants
Go constants are unique because they can be **untyped**. 
*   An untyped constant has no fixed type, but carries a high-precision value (up to 256 bits).
*   It only assumes a concrete type when assigned to a variable or used in an expression requiring a type.
```go
const Pi = 3.14159265358979323846 // Untyped float constant
var f32 float32 = Pi // Automatically converted to float32
var f64 float64 = Pi // Automatically converted to float64 (no warning)
```

### The `iota` Constant Generator
Inside a constant block, the pre-declared identifier `iota` acts as an integer generator starting at `0`, incrementing by 1 for each line:
```go
const (
    Low = iota // 0
    Medium     // 1 (implicitly = iota)
    High       // 2 (implicitly = iota)
)
```
You can use `iota` with bitwise operations to represent flags:
```go
const (
    Read   = 1 << iota // 1 << 0 = 1
    Write              // 1 << 1 = 2
    Execute            // 1 << 2 = 4
)
```

---

## 🔍 Deep Dive 4: Zero Values

If a variable is declared but not initialized, Go automatically initializes its memory to its **zero value**.

| Type Family | Zero Value |
| :--- | :--- |
| Integer and Float | `0` or `0.0` |
| Boolean | `false` |
| String | `""` (empty string) |
| Pointer (`*T`) | `nil` |
| Interface (`interface{}`) | `nil` |
| Reference Types (`slice`, `map`, `channel`, `func`) | `nil` |
| Struct | All nested fields recursively set to their zero values |

---

## ⚠️ Common Gotchas

1.  **Unused Constants**: Unlike variables, the Go compiler **does not** throw an error for unused constants. This is because constants are evaluated at compile-time and do not generate runtime allocations.
2.  **Short-Declaration Shadowing**: If you declare a variable at the package level, and then redeclare it inside a function block using `:=`, it creates a new variable that **shadows** the package variable inside that block. This can lead to subtle bugs where you think you are updating a global variable but are actually modifying a local copy.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement the function to return the correct types, and verify compilation.
```bash
go run .
```
