# Step 1.5: Pointers Demystified 📍

Pointers let you refer directly to memory addresses. 

## Analogy
Instead of carrying your entire house to your friend, you write down your address on a piece of paper. The paper with the address is a **pointer**.
- `&` operator gets the address.
- `*` operator dereferences (gets the value at that address).

## Key Concepts
1. **Pass-by-value**: By default, Go copies arguments. If you want to modify a variable inside a function, you must pass its pointer.
2. **`nil` pointer**: A pointer that points to nothing.

## How to Run
```bash
go run main.go
```
