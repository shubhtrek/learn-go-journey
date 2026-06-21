# Step 2.1: Slices vs. Arrays (Memory Headers & Dynamic Allocation) 🍕

This step covers the mechanics of fixed-size arrays and dynamic slices, details their memory representation, and explains the runtime behavior of backing array allocation.

Official documentation:
*   [Go Spec: Array Types](https://golang.org/ref/spec#Array_types)
*   [Go Spec: Slice Types](https://golang.org/ref/spec#Slice_types)
*   [Go Blog: Go Slices: usage and internals](https://go.dev/blog/slices-intro)
*   [Effective Go: Slices](https://go.dev/doc/effective_go#slices)

---

## 🔍 Deep Dive 1: Fixed-Size Arrays

An **Array** is a numbered sequence of elements of a single type, called the element type. The number of elements is the length of the array and is **part of the array's type**.
```go
var a [5]int  // Type is [5]int
var b [10]int // Type is [10]int
// a = b      // ❌ Compile error: cannot use b (type [10]int) as type [5]int in assignment
```
*   **Compile-time Fixed Size**: The size must be a constant expression that evaluates to a non-negative integer.
*   **Value Type**: In Go, arrays are value types, not reference types. When you assign an array to another, or pass it to a function, Go **copies the entire array** in memory.

---

## 🔍 Deep Dive 2: Slice Headers and Memory Layout

A **Slice** is a dynamically sized, flexible view into the elements of an array. Unlike arrays, the length of a slice is not part of its type (written as `[]T`).

### The Slice Header
Under the hood, a slice is represented as a runtime struct called the **Slice Header** (previously defined in `reflect.SliceHeader`). It occupies exactly **24 bytes** of memory on a 64-bit architecture:

```text
+-----------------------+-----------------------+-----------------------+
|  Backing Array Pointer|      Length (len)     |    Capacity (cap)     |
|       (8 bytes)       |       (8 bytes)       |       (8 bytes)       |
+-----------------------+-----------------------+-----------------------+
```

1.  **Pointer**: The memory address of the first element of the slice (which may not be the first element of the underlying backing array).
2.  **Length (`len`)**: The number of elements currently in the slice. Retrieved via `len(s)`.
3.  **Capacity (`cap`)**: The total number of elements in the underlying backing array, starting from the slice's first element. Retrieved via `cap(s)`.

---

## 🔍 Deep Dive 3: The `append()` Capacity Growth Algorithm

When you append elements to a slice using the built-in `append(slice, elements...)` function:

1.  **Under Capacity (`len + n <= cap`)**: Go simply writes the new values into the existing backing array, updates the slice header's length field, and returns the modified slice header.
2.  **Out of Capacity (`len + n > cap`)**: Go allocates a **new backing array** in the runtime heap, copies the existing elements, appends the new values, and returns a new slice header pointing to this new location.

### Dynamic Allocation Rules
Prior to Go 1.18, the capacity allocation algorithm simply doubled the size of the slice up to 1024 elements, and then scaled by `1.25x` for larger slices.
In modern Go (Go 1.18+), the growth algorithm transitions smoothly to prevent abrupt memory spikes:
*   If the required capacity is larger than double the old capacity, the new capacity is set to the required capacity.
*   Otherwise, if the old capacity is less than **256**, the new capacity is doubled.
*   If the old capacity is 256 or more, the new capacity is calculated as: `newcap = oldcap + (oldcap + 3*256)/4` (which converges toward a `1.25x` growth rate gradually).

---

## 🔍 Deep Dive 4: Slicing Operations & Three-Index Slicing

Slicing creates a new slice header pointing to the same backing array.

### 1. Two-index Slicing (`slice[low:high]`)
Creates a view starting at `low` up to (but excluding) `high`.
*   `Length` = `high - low`
*   `Capacity` = `cap(slice) - low`
*   **Warning**: Modifying elements in `slice[low:high]` mutates the shared backing array, affecting the parent slice.

### 2. Three-index Slicing (`slice[low:high:max]`)
Controls the capacity of the resulting slice, preventing subsequent `append()` operations from writing over elements of the parent backing array:
*   `Length` = `high - low`
*   `Capacity` = `max - low`
*   **Constraint**: `low <= high <= max <= cap(slice)`

---

## ⚠️ Common Gotchas & Memory Leaks

### 1. Shared Backing Array Side-effects
```go
parent := []int{1, 2, 3, 4}
sub := parent[0:2] // sub is [1, 2] with cap 4
sub = append(sub, 99) // Mutates index 2 of parent!
fmt.Println(parent) // prints [1, 2, 99, 4]
```
To avoid this, you can copy the slice explicitly to a new memory segment using the built-in `copy` function:
```go
dest := make([]int, len(sub))
copy(dest, sub) // dest now points to an independent backing array
```

### 2. Garbage Collection Pinning (Memory Leak)
If a function reads a massive file into a huge backing slice, extracts a small subset (e.g. 2 bytes), and returns that subset slice:
```go
func getSmallData() []byte {
    hugeData := readHugeFile() // e.g. 100MB slice
    return hugeData[0:2] // ⚠️ Memory Leak! The 100MB backing array cannot be GC'ed because the returned 2-byte slice still points to it.
}
```
**Fix**: Allocate a new small slice, copy the data, and let the massive slice go out of scope to allow garbage collection:
```go
func getSmallData() []byte {
    hugeData := readHugeFile()
    result := make([]byte, 2)
    copy(result, hugeData[0:2])
    return result
}
```

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement the slice filtering challenges. Run:
```bash
go run .
```
Verify the output structure.
