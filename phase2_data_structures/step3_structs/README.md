# Step 2.3: Structs & Methods 🏗️

Go uses structs and methods to model data and attach behaviors, giving you object-oriented benefits without inheritance bloat.

---

## 🔍 Deep Dive: Value vs Pointer Receivers

A **method** is a function with a special receiver argument placed before the function name:
```go
func (r ReceiverType) MethodName() { ... }
```

### Value Receivers (`func (a Account) GetBalance()`)
*   Go copies the entire struct to pass it to the method.
*   Modifications to fields inside the method affect only the copy, not the original struct.
*   Used for small structs or read-only methods.

### Pointer Receivers (`func (a *Account) Deposit(...)`)
*   Go passes the memory address of the struct.
*   Modifications directly affect the original struct.
*   Highly efficient because it avoids copying large structs.
*   **Rule of thumb**: If any method on a struct requires a pointer receiver, *all* methods on that struct should have pointer receivers to ensure consistency.

---

## ⚠️ Common Gotchas
1.  **Modifying state in Value Receivers**: If you forget the `*` in the receiver type, your method will modify a temporary copy and your actual object state won't update, causing confusing bugs.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go).
1. Implement the `Deposit` and `Withdraw` methods on the `BankAccount` struct.
2. `Withdraw` should check if the account has enough funds and return an error if it doesn't.
3. Test your struct and methods in `main.go`.
4. Tell me when you are ready for a review!

---

## 🎉 Shubham's Fun Corner 🎉

### 🍱 The "Special Thali" Analogy
A struct is like ordering a special Indian Thali. Instead of ordering rice, dal, and paneer separately, you get a single pre-defined plate containing:
*   Rice (`string`)
*   Dal (`float64`)
*   Paneer (`int`)
*   Samosa (`bool`)
It groups different items into a single delicious package!

### ☕ Chai Break Thought
Pointer receivers let you change the food inside the Thali. Value receivers let you look at the food, but you can't eat or modify it. Choose your receivers wisely!


### 👑 Marathi Swag: The "Special Maharaja Thali"
*   A struct is like ordering a **Maharaja Thali**.
*   Instead of ordering Chapati, Paneer, Rice, and Gulab Jamun separately, you get a single plate containing everything:
    ```go
    type Thali struct {
        Chapati int
        Paneer  string
        Sweet   bool
    }
    ```
*   **Pointer receiver (`*Thali`)**: Lets you actually eat or modify the items inside the plate.
*   **Value receiver (`Thali`)**: You are just looking at a photo of the Thali. You can't change it!
