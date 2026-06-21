# Step 2.2: Maps (Hash Map Internals & Concurrency Limitations) 🗺️

This step covers the mechanics of Go's built-in map type, detailing how the runtime structures hash maps, how key-value lookups operate, and why maps are unsafe for concurrent operations.

Official documentation:
*   [Go Spec: Map Types](https://golang.org/ref/spec#Map_types)
*   [Go Blog: Go maps in action](https://go.dev/blog/maps)
*   [Effective Go: Maps](https://go.dev/doc/effective_go#maps)

---

## 🔍 Deep Dive 1: Map Under-the-Hood Architecture

In Go, a map is a reference to a runtime struct called `hmap` (defined in `runtime/map.go`). 

```text
hmap struct:
+------------------+------------------+-------------------+
|      count       |      flags       |         B         |
|  (number of keys)| (concurrency flg)| (log_2 of buckets)|
+------------------+------------------+-------------------+
|    hash0 seed    |   buckets ptr    |   oldbuckets ptr  |
|  (hash generator)| (current storage)|  (during growth)  |
+------------------+------------------+-------------------+
```

### 1. Buckets & Memory Layout
*   A map contains a pointer to an array of **buckets** (`bmap` struct).
*   Each bucket holds up to **8 key-value pairs**.
*   The compiler calculates the hash value of the key (using algorithm specific to the key type, e.g. AES hashing).
*   The low-order bits of the hash are used to select which bucket the key belongs to. The high-order bits (Top Hash or `tophash`) are stored inside the bucket to quickly identify keys during lookups.

### 2. Overflow Buckets
If more than 8 keys hash to the same bucket (hash collision), Go allocates an **overflow bucket** and links it to the original bucket, forming a linked list. If overflow chains grow too long, or the load factor exceeds 6.5, Go triggers a **map evacuation (growth)**, allocating double the buckets and moving data incrementally to prevent latency spikes.

---

## 🔍 Deep Dive 2: Map Operations & Syntax

### 1. Allocation
Always initialize maps using `make()` or map literals. Declaring a map without initializing creates a `nil` map:
```go
var m map[string]int // m is nil
// m["key"] = 1      // ❌ RUNTIME PANIC: assignment to entry in nil map

m = make(map[string]int) // ✅ Correct: allocates hmap
m = make(map[string]int, 100) // ✅ Correct: pre-allocates space for 100 keys (saves resizing overhead)
```

### 2. The Comma-Ok Existence Check
If a key does not exist in a map, lookup returns the **zero value** of the value type. To distinguish between a key set to its zero value versus a key that doesn't exist, use the "comma-ok" idiom:
```go
value, ok := m["unknown_key"]
if !ok {
    fmt.Println("Key does not exist!")
}
```

### 3. Deleting Keys
Delete keys using the built-in `delete(map, key)` function. If the key doesn't exist, `delete` does nothing and does not crash.
```go
delete(m, "key")
```

---

## 🔍 Deep Dive 3: Map Non-Addressability

You cannot take the address of a map element:
```go
m := map[string]int{"a": 1}
// ptr := &m["a"] // ❌ Compile error: cannot take the address of m["a"]
```
### Why?
Because maps grow and evacuate data dynamically. If the runtime redistributes keys into new buckets during a map growth operation, the elements are moved to different memory addresses. Allowing pointers to map values would result in dangling pointers and memory unsafety.

---

## 🔍 Deep Dive 4: Concurrency Limitations

**Go maps are NOT thread-safe.** 

If one goroutine is writing to a map while another goroutine is reading or writing to the same map concurrently without synchronization, the Go runtime detects this and throws an unrecoverable crash:
```text
fatal error: concurrent map read and map write
```
This is not a standard panic; it cannot be recovered via `recover()`. It immediately terminates the program to prevent memory corruption.

### Solutions for Concurrency
1.  **Mutex Wrapper**: Wrap the map in a struct along with a `sync.Mutex` or `sync.RWMutex`:
    ```go
    type SafeMap struct {
        sync.RWMutex
        data map[string]int
    }
    ```
2.  **`sync.Map`**: Use the standard library's `sync.Map`, which is optimized for specific workloads (e.g. read-heavy, disjoint key sets).

---

## ⚠️ Common Gotchas

1.  **Map Iteration Randomization**: Iterating over a map via `for range` returns keys in a randomized order. This is a deliberate design decision in Go's runtime: the starting bucket is selected at random using a random seed to prevent developers from relying on a specific key ordering.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement map lookup and deletion operations. Run:
```bash
go run .
```
Verify compilation.
