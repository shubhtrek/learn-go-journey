# Step 2.1: Slices vs Arrays 🍕

Slices are the primary data sequence type in Go. Let's see how they work under the hood.

---

## 🔍 Deep Dive: The Slice Header & Dynamic Growth

An **Array** has a fixed size defined at compile-time: `var arr [5]int`. You cannot resize it.
A **Slice** is a dynamic view into a backing array. 

### The Slice Header Structure
Under the hood, a slice is a struct (Slice Header) that occupies only 24 bytes in memory:
1.  **Pointer**: Address of the first element in the backing array.
2.  **Length (`len`)**: Number of elements currently in the slice.
3.  **Capacity (`cap`)**: Maximum elements the slice can hold before resizing the backing array.

### How `append()` works:
When you call `slice = append(slice, value)`:
1. If `len < cap`, Go writes the value into the existing backing array and increases `len` by 1.
2. If `len == cap`, Go creates a new backing array (usually double the size), copies old elements, appends the new value, updates the slice header's pointer, and updates `cap`.

---

## ⚠️ Common Gotchas
1.  **Backing Array Sharing**: If you slice a slice (`sub := slice[1:3]`), they still share the same backing array. Modifying `sub[0]` will change the value in `slice[1]`!
2.  **Memory Leaks**: Keeping a small slice of a massive array in memory prevents the whole backing array from being garbage collected. Use `copy()` to create a clean, independent slice if needed.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go).
1. Implement `FilterEven` to filter even numbers from a slice.
2. Write code in `main.go` to test your slice function.
3. Tell me when you are ready for a review!

---

## 🎉 Shubham's Fun Corner 🎉

### 👖 The "Elastic Track Pants" Analogy
*   **Arrays** are like formal wedding suits: fixed size. If you gain even a little weight (append elements), the suit rips (compile error).
*   **Slices** are like elastic track pants: dynamic. They stretch automatically as you eat more samosas. Under the hood, Go silently swaps them for a larger pair (backing array) when you exceed capacity.

### ☕ Chai Break Thought
When you slice a slice, you are just looking at a section of the same cloth. Don't be surprised if modifying the sub-slice colors the main slice!
