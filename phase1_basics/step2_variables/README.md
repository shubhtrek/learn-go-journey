# Step 1.2: Variables, Types, & Zero Values 📦

Go is statically typed, meaning once a variable is a certain type (like `string` or `int`), it cannot change.

## Analogy
You can't pour hot tea (`chai`) into a paper envelope. Go makes sure the container type matches the data type!

## Key Concepts
1. **Explicit Declaration (`var`)**: `var name string = "Shubham"`
2. **Short Assignment (`:=`)**: `age := 20` (Go infers the type automatically, only works inside functions).
3. **Zero Values**: Uninitialized variables get a default value (`0` for ints, `""` for strings, `false` for bools). No more random `undefined` crashes!

## How to Run
```bash
go run main.go
```
