# Step 2.3: Structs, Methods & Memory Alignment 🏗️

This step covers how Go groups data using structures, attaches behavior using methods, and optimizes struct memory footprint in physical RAM through byte alignment and padding.

Official documentation:
*   [Go Spec: Struct Types](https://golang.org/ref/spec#Struct_types)
*   [Go Spec: Method Declarations](https://golang.org/ref/spec#Method_declarations)
*   [Go Doc: Effective Go: Allocation with new](https://go.dev/doc/effective_go#allocation_new)

---

## 🔍 Deep Dive 1: Memory Alignment and Struct Padding

A struct is a sequence of named elements, called fields. While a struct seems like a simple list of fields, the compiler arranges them in memory to optimize CPU access speeds. This is known as **Memory Alignment**.

### 1. CPU Word Size & Alignment Rules
Modern CPUs read memory in "words" rather than single bytes (typically 8 bytes on a 64-bit CPU). To read memory efficiently:
*   A variable of size $n$ bytes must be stored at a memory address that is a multiple of $n$ (alignment boundary).
*   For example, an `int64` (8 bytes) must be stored at an address divisible by 8. A `bool` (1 byte) can be stored at any address.

### 2. Struct Padding Example
If you declare fields in an sub-optimal order, the compiler injects **padding bytes** (wasted space) to align subsequent fields:

```go
type BadStruct struct {
    a bool   // 1 byte
             // 7 bytes padding (to align next int64 field)
    b int64  // 8 bytes
    c bool   // 1 byte
             // 7 bytes padding (to align overall struct to 8-byte boundary)
}
// Size: 24 bytes!
```

By grouping fields from largest to smallest, you minimize padding:
```go
type GoodStruct struct {
    b int64  // 8 bytes
    a bool   // 1 byte
    c bool   // 1 byte
             // 6 bytes padding (to align overall struct to 8-byte boundary)
}
// Size: 16 bytes! (Saved 8 bytes)
```
You can inspect structural sizes using `unsafe.Sizeof()`.

---

## 🔍 Deep Dive 2: Value Receivers vs. Pointer Receivers

Go allows you to define methods on struct types. A method is simply a function with an explicit **receiver** argument:
```go
func (s MyStruct) ValueMethod() {}
func (s *MyStruct) PointerMethod() {}
```

### Choosing the Receiver Type

| Aspect | Value Receiver `(s MyStruct)` | Pointer Receiver `(s *MyStruct)` |
| :--- | :--- | :--- |
| **Mutation** | Operates on a **copy** of the struct. Modifications do not affect the caller. | Operates on the **original** memory. Modifying fields updates the caller's struct. |
| **Performance** | Copies the entire struct on every call. Can be slow if the struct contains large arrays/fields. | Copies only the pointer address (8 bytes). Faster for large structs. |
| **Method Sets** | Can be called on both values and pointers. | Can be called on pointers. Can be called on addressable values (Go automatically inserts `&` under the hood). |

### The Method Set Rules
The Go specification defines which methods are available on which types:
*   For a value of type `T`: The method set contains only methods with receiver `T`.
*   For a pointer of type `*T`: The method set contains methods with receivers `T` and `*T`.
*   **Interface Satisfaction**: If an interface requires a pointer receiver method, a value of type `T` **cannot** satisfy the interface. You must pass a pointer of type `*T` to the interface.

---

## 🔍 Deep Dive 3: Composition over Inheritance (Struct Embedding)

Go does not have classes or classical inheritance. Instead, it promotes reuse through **Composition** via anonymous struct embedding:

```go
type Engine struct {
    Horsepower int
}

type Car struct {
    Engine // Anonymous field: Engine is embedded in Car
    Brand  string
}
```

### Field Promotion
When a struct is embedded anonymously:
*   The fields and methods of the embedded struct are **promoted** to the parent struct.
*   You can access embedded fields directly: `myCar.Horsepower` instead of `myCar.Engine.Horsepower`.
*   If name collisions occur, you can resolve them by explicitly referencing the embedded type name: `myCar.Engine.Horsepower`.

---

## 🔍 Deep Dive 4: Constructor Patterns

Go does not have a formal `constructor` keyword. Instead, the idiomatic pattern is to write a helper function (typically prefixed with `New`) that returns an initialized instance or pointer to the struct:
```go
type Client struct {
    host string
    port int
}

func NewClient(host string, port int) *Client {
    return &Client{
        host: host,
        port: port,
    }
}
```

---

## ⚠️ Common Gotchas

1.  **Empty Structs (`struct{}`)**: An empty struct contains no fields and occupies exactly **0 bytes** of memory. It is highly useful as a placeholder (e.g. signaling channel events `chan struct{}` or map sets `map[string]struct{}`) because it generates zero allocations.
2.  **Pointer Receiver on Nil Structs**: A method with a pointer receiver can still be called if the receiver is `nil`. It is your responsibility to handle null checks inside the method to prevent nil-pointer panics when accessing fields:
    ```go
    func (c *Client) Connect() {
        if c == nil {
            return // Prevent panic!
        }
        // ...
    }
    ```

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement struct structures, constructor functions, and methods. Run:
```bash
go run .
```
Verify compilation.
