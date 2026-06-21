# Step 2.1: Slices vs Arrays 🍕

In Go, arrays have a fixed size. Slices are dynamic wrappers around arrays that can grow.

## Analogy
An array is a bench with exactly 3 seats. A slice is an elastic sofa that expands as more people arrive.

## Key Concepts
1. **Slice growth**: Go automatically allocates a larger backing array when a slice exceeds its capacity.
2. **`append()`**: Built-in function to add elements to a slice.
3. **`make()`**: Used to initialize slices with predefined length and capacity.

## How to Run
```bash
go run main.go
```
