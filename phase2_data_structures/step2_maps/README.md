# Step 2.2: Maps (Key-Value Pairs) 🗺️

Maps are fast lookup tables. Let's study how they work and how to safely access keys.

---

## 🔍 Deep Dive: How Maps Work & The Comma-Ok Idiom

A map is a reference type pointing to a hash table.

### Initialization
You should initialize maps using `make()`:
```go
var m map[string]int // m is nil! Writing to it will CRASH.
m = make(map[string]int) // Now it is allocated and ready.
```

### Accessing keys (The Comma-Ok Idiom)
If a key is not present in a map, accessing it returns the zero value of the value type:
```go
score := scores["unknown"] // returns 0
```
To distinguish between a key that has a value of `0` and a key that doesn't exist, use the **comma-ok** idiom:
```go
score, ok := scores["unknown"]
if ok {
    // Key exists
} else {
    // Key does not exist
}
```

---

## ⚠️ Common Gotchas
1.  **Writing to a `nil` map**: Declaring a map variable `var m map[string]int` without `make()` creates a `nil` map. Reading from it is safe, but writing `m["key"] = 10` causes a runtime panic!
2.  **Concurrency unsafe**: Go maps are *not* thread-safe. Multiple goroutines reading and writing to the same map concurrently will crash the app. You must use a Mutex (which we learn in Phase 4).

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go).
1. Implement `GetPrice` to look up item prices in a menu map using the comma-ok idiom.
2. Write code in `main.go` to test this map lookup.
3. Let me know when you've written the solution!
