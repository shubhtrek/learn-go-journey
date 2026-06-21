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

---

## 🎉 Shubham's Fun Corner 🎉

### 🚗 The "Potholes on Indian Roads" Analogy
In other languages, people write code and hope for the best, letting exceptions crash the app later. 
In Go, handling errors is like driving on city roads: you look out for potholes every few meters (`if err != nil`). It requires attention, but it guarantees you and your car arrive safely!

### ☕ Chai Break Thought
Errors are not exceptions; they are just variables. Treat them with respect, don't ignore them, and keep smiling!


### 👑 Marathi Swag: The "Mumbai Traffic" Rule
*   Go does not have try-catch blocks. 
*   Writing Go code is like driving on a highway full of potholes. You don't close your eyes and pray (try-catch). 
*   You check for every pothole carefully: **`if err != nil { // take bypass }`**.
*   It might look like extra typing, but it prevents your car (program) from crashing!
