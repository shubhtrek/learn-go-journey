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
