# Step 2.5: Error Handling (No Try-Catch) ⚠️

Go treats errors as first-class values. Let's see how error handling works under the hood.

---

## 🔍 Deep Dive: The Error Interface & Explicit Checks

In Go, there is no exception throwing mechanism (`throw/try/catch`). Instead, errors are returned as values alongside the successful result.

### The `error` Interface
The built-in `error` type is a simple interface:
```go
type error interface {
    Error() string
}
```
Any type that implements the `Error() string` method can be used as an error.

### Creating errors:
*   `errors.New("message")`: For simple, static error messages.
*   `fmt.Errorf("failed to load user %d", id)`: To format messages with variables.

---

## ⚠️ Common Gotchas
1.  **Ignoring errors**: Never ignore errors using `_`. Always check `if err != nil` and handle it (log it, return it, or recover from it).
2.  **Nil Interface Trap**: An interface is only `nil` if both its type and its value are `nil`. If you return a custom error struct pointer that is nil, the interface checking `err != nil` might evaluate to `true`! (Stick to returning the standard `error` type directly).

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go).
1. Implement `SafeDivide` to perform division and return a custom error if divisor is `0`.
2. Test your implementation in `main.go` and handle the error gracefully by printing a warning.
3. Message me for review!
