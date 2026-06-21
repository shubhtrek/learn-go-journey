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

---

## 🎉 Shubham's Fun Corner 🎉

### 👵 The "Grandmother's Memory" Analogy
Maps are like a contact list in your grandmother's head. You ask: *"Dadi, Ramesh?"* She instantly responds: *"Sharma ji ka beta, class topper, lives in flat 4B."* 
She doesn't search through a notebook from page 1. The search time is instantaneous—**O(1)**.

### ☕ Chai Break Thought
Remember the comma-ok check! It's like asking: *"Dadi, does Ramesh exist?"* before you try to call him. If you don't check, Go might return a silent `0` and you'll call the wrong person!


### 👑 Marathi Swag: The "Kirana Shop Uncle"
*   Maps are like the local Kirana shop uncle. You ask: *"Uncle, Pohe kuthe aahet?"* (Where is the Pohe?).
*   He doesn't search the whole shop. He instantly points to rack 2. **O(1) search speed, bhava!**
*   But remember the comma-ok check! Always check: `value, exist := shop["Pohe"]`. If `exist` is false, don't try to buy it, otherwise you'll get a null panic!
