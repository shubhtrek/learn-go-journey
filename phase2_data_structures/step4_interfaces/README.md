# Step 2.4: Interfaces (Implicit Contracts & Runtime Representations) 🔌

This step covers how Go defines abstract behaviors, how the compiler represents interfaces under the hood, and how type assertions and type switches work.

Official documentation:
*   [Go Spec: Interface Types](https://golang.org/ref/spec#Interface_types)
*   [Go Spec: Type Assertions](https://golang.org/ref/spec#Type_assertions)
*   [Effective Go: Interfaces](https://go.dev/doc/effective_go#interfaces)

---

## 🔍 Deep Dive 1: Duck Typing and Implicit Satisfaction

An interface type specifies a **method set** called its interface. A variable of interface type can store any value that implements these methods.
*   **Implicit satisfaction**: Unlike Java or C#, Go structs do not use an `implements` keyword. If a struct implements all methods defined in an interface, the compiler automatically satisfies the interface.
*   **Benefits**: Decouples code completely. A consumer can define an interface representing the behavior it needs, and any third-party struct can be passed in without modification.

---

## 🔍 Deep Dive 2: Interface Internals (`iface` vs. `eface`)

At the runtime level, interfaces are not simple pointers. They are represented by two distinct structures (defined in `runtime/runtime2.go`):

### 1. `iface` (Non-empty Interface)
Used for interfaces that declare methods (e.g. `io.Reader`). It occupies 16 bytes:
*   **`tab` (pointer to an `itab` struct)**: Contains metadata about the interface type, the concrete type description, and a table of function pointers (the actual method implementations for that type).
*   **`data` (unsafe pointer)**: Points to the actual value of the concrete type stored in the heap.

```text
iface struct:
+-------------------+-------------------+
|     itab ptr      |     data ptr      |
|  (type & methods) | (concrete value)  |
+-------------------+-------------------+
```

### 2. `eface` (Empty Interface: `interface{}` or `any`)
Used for the empty interface (`interface{}` or `any`). Because it has no methods, it doesn't need an `itab` table. It occupies 16 bytes:
*   **`_type` (pointer to a `_type` struct)**: Holds representation metadata about the concrete type (its size, hash value, etc.).
*   **`data` (unsafe pointer)**: Points to the concrete value.

```text
eface struct:
+-------------------+-------------------+
|     _type ptr     |     data ptr      |
| (metadata only)   | (concrete value)  |
+-------------------+-------------------+
```

---

## 🔍 Deep Dive 3: Type Assertions and Type Switches

To extract the concrete value stored inside an interface, Go uses type assertions and type switches.

### 1. Type Assertions
A type assertion expression checks if the interface value holds a specific concrete type:
```go
var w io.Writer = os.Stdout
f, ok := w.(*os.File) // Checks if w holds a pointer to os.File
if ok {
    // f is now of type *os.File
    fmt.Println(f.Name())
}
```
*   **Syntax**: `value, ok := interfaceVariable.(TargetType)`
*   **Safety**: If you omit the `ok` variable (e.g. `f := w.(*os.File)`) and the assertion fails, the program will **runtime panic**. Always use the comma-ok form for safety.

### 2. Type Switches
A type switch compares the interface type against multiple concrete types:
```go
func inspect(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %q\n", v)
    default:
        fmt.Println("Unknown type")
    }
}
```

---

## ⚠️ The Nil-Interface Trap (Crucial Gotcha)

An interface is only considered `nil` if **both** its type descriptor (`tab` or `_type`) and its value pointer (`data`) are `nil`. 

If you assign a `nil` pointer of a concrete type to an interface, the interface variable itself is **not nil** because its type descriptor is set:

```go
var buf *bytes.Buffer // buf is nil
var w io.Writer = buf // w holds (*bytes.Buffer, nil)

fmt.Println(w == nil) // ❌ Prints FALSE! 
```

### Why is this dangerous?
If you pass `w` to a function that checks for `w != nil` before writing, the check will pass. However, trying to invoke methods on it will trigger a nil-pointer panic inside the method because the underlying data pointer is `nil`.

**Prevention**: Always return interfaces explicitly as `nil` rather than returning concrete pointer variables that happen to be `nil`.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement interface assignments, assertions, and verification. Run:
```bash
go run .
```
Verify compilation.
