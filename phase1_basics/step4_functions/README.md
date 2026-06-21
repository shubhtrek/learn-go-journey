# Step 1.4: Functions & Multiple Returns 📞

Functions are the building blocks of Go. Let's study how parameters are passed and how functions return data.

---

## 🔍 Deep Dive: Value Semantics & Named Returns

### Pass-by-Value (Copying)
Go is strictly **pass-by-value**. When you pass an argument to a function, Go creates a copy of that data in memory and passes the copy to the function. Any modifications inside the function do not affect the original variable (unless you pass a pointer, which we will learn next).

### Multiple Return Values
Go functions can return multiple values. This is heavily used for error handling, returning the result and the error:
```go
func fetchRecord(id int) (Record, error) { ... }
```

### Named Return Values
You can name the return variables in the function signature. They are treated as variables defined at the top of the function and are initialized to their zero values:
```go
func getCoords() (x int, y int) {
    x = 10
    y = 20
    return // "Naked return" - returns x and y automatically
}
```
*Tip: Avoid naked returns in long functions as it degrades code readability.*

---

## ⚠️ Common Gotchas
1.  **Ignoring Return Values**: If a function returns multiple values, you must handle all of them. Use the blank identifier `_` to discard values you don't need:
    ```go
    sum, _ := calculate(10, 5) // Discards the second return value
    ```

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go).
1. Implement `CalculateRectangle` to return both the Area and Perimeter of a rectangle using named return values.
2. Test it by calling it from `main.go`.
3. Let me know when you are ready for a review!

---

## 🎉 Shubham's Fun Corner 🎉

### 👦 The "Younger Brother" Analogy
Multiple return values are like sending your younger brother to the local shop. You tell him to get milk. He returns with **milk** AND the **change**. If the shop is closed, he returns with **no milk** and an **excuse** (error). 
In Go, functions do exactly this: `result, err := doSomething()`

### ☕ Chai Break Thought
You can name your return values in Go. It's like naming your children before they are born—very organized!
