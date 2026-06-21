# Step 2.4: Interfaces (Go's Secret Weapon) 🔌

Interfaces decouple what an object *does* from what it *is*. Let's understand implicit interfaces.

---

## 🔍 Deep Dive: Implicit Interfaces & Polymorphism

In languages like Java or C#, you must explicitly declare that a class implements an interface:
```java
class Dog implements Speaker { ... }
```
In Go, interfaces are implemented **implicitly**:
If a struct implements all the methods defined in an interface, Go automatically considers that struct as implementing the interface. No keyword required!

### The Empty Interface `interface{}` (or `any`)
An interface with zero methods is implemented by *every* type. In Go 1.18+, `any` was introduced as an alias for `interface{}`:
```go
func PrintAnything(val any) { ... }
```

---

## ⚠️ Common Gotchas
1.  **Pointer vs Value Receivers in Interfaces**: If a struct implements an interface method using a pointer receiver, you *must* pass a pointer of that struct to the interface variable.
    ```go
    var s Speaker
    s = Dog{}  // ❌ COMPILE ERROR if Speak() has a pointer receiver (*Dog)
    s = &Dog{} // ✅ Correct
    ```

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go).
1. Define the `Area()` method for `Rectangle` and `Circle` structs so they implicitly satisfy the `Shape` interface.
2. In `main.go`, write a function `PrintArea(s Shape)` and pass both structs to it to print their areas.
3. Message me for manual review!
