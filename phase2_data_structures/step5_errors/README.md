# Step 2.5: Error Handling (Explicit Verification & Wrapped Errors) ⚠️

This step covers how Go represents errors as standard values, implements custom error types, wraps errors to preserve context, and checks them idiomatic way.

Official documentation:
*   [Go Spec: Errors](https://golang.org/ref/spec#Errors)
*   [Go Blog: Go 1.13 Errors](https://go.dev/blog/go1.13-errors)
*   [Effective Go: Errors](https://go.dev/doc/effective_go#errors)

---

## 🔍 Deep Dive 1: Errors as Values

Unlike languages that use exceptions (like Java, Python, or JS), Go does not have a `try-catch` mechanism. Instead, errors are treated as **normal values** returned explicitly from functions.

The built-in `error` type is a simple interface:
```go
type error interface {
    Error() string
}
```
Any type that implements the `Error() string` method satisfies the `error` interface.

### Benefits of "Errors as Values"
1.  **Linear Control Flow**: Control flow remains clear and predictable. You do not jump blocks unexpectedly.
2.  **Explicit Handling**: Forces developers to handle failures where they occur rather than letting them propagate unhandled up the stack.

---

## 🔍 Deep Dive 2: Sentinel Errors vs. Custom Errors

### 1. Sentinel Errors
Sentinel errors are pre-declared package-level error variables indicating a specific state:
```go
var ErrNotFound = errors.New("item not found")
```
*   You check for sentinel errors by performing a direct equality comparison: `if err == ErrNotFound`.
*   **Limitation**: Sentinel errors cannot carry dynamic context (like a user ID or database query).

### 2. Custom Error Structs
When you need to attach extra diagnostic information to an error, implement a custom struct:
```go
type QueryError struct {
    Query string
    Err   error
}

func (e *QueryError) Error() string {
    return fmt.Sprintf("query %q failed: %v", e.Query, e.Err)
}
```

---

## 🔍 Deep Dive 3: Error Wrapping and Unwrapping

Go 1.13 introduced native support for wrapping errors to preserve call-stack context without losing the original error type.

### 1. Wrapping Errors (`%w`)
To wrap an error, use `fmt.Errorf` with the `%w` format verb:
```go
func ReadConfig() error {
    err := openFile()
    if err != nil {
        return fmt.Errorf("failed to read config: %w", err) // Wraps err
    }
    return nil
}
```

### 2. Checking Wrapped Errors (`errors.Is`)
Use `errors.Is` to check if a wrapped error chain contains a specific sentinel error:
```go
err := ReadConfig()
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File does not exist!") // True even if wrapped!
}
```

### 3. Extracting Wrapped Errors (`errors.As`)
Use `errors.As` to see if any error in the wrapped chain matches a custom error type, and if so, bind it to a variable:
```go
var queryErr *QueryError
if errors.As(err, &queryErr) {
    fmt.Println("Failed query was:", queryErr.Query) // Binds and extracts the struct!
}
```

---

## ⚠️ Common Gotchas

1.  **Errors as Value-Type Nil Trap**: If you declare a custom pointer-based error type and return it directly from a function, it can trigger the nil-interface trap:
    ```go
    func check() error {
        var err *QueryError = nil
        return err // ⚠️ Returns an interface containing (*QueryError, nil) which is NOT nil!
    }
    ```
    **Prevention**: Always return `nil` explicitly rather than returning a typed pointer that is `nil`.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement custom error wrapping and checks using `errors.Is` and `errors.As`. Run:
```bash
go run .
```
Verify compilation.
