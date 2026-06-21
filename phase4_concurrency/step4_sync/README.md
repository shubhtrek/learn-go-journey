# Step 4.4: The Sync Package (Locks, WaitGroups & Memory Barriers) 🔒

This step covers low-level synchronization primitives, CPU-level atomic instructions, memory synchronization boundaries, and detecting data race conditions.

Official documentation:
*   [Go Package: sync](https://pkg.go.dev/sync)
*   [Go Package: sync/atomic](https://pkg.go.dev/sync/atomic)
*   [Go Blog: Introducing the Go Race Detector](https://go.dev/blog/race-detector)

---

## 🔍 Deep Dive 1: `sync.WaitGroup` Internals

A `sync.WaitGroup` is used to block execution until a set of goroutines finishes executing.

```text
WaitGroup Internal State (64-bit atomic counter):
+-------------------------------+-------------------------------+
|      Counter (32 bits)        |      Waiter Count (32 bits)   |
| (active goroutines running)   |  (goroutines blocked in Wait) |
+-------------------------------+-------------------------------+
```

*   `Add(n)`: Increments the active counter by $n$.
*   `Done()`: Decrements the active counter by 1. Calls `runtime_Semasignals` if the counter reaches 0 to unblock waiting goroutines.
*   `Wait()`: Increments the waiter count and parks the calling goroutine until the counter drops to 0.

### Rules of WaitGroup
1.  **Do not copy**: A WaitGroup contains internal state that cannot be copied. If passed to a function, you must pass its **address** (`*sync.WaitGroup`), otherwise it will copy the state by value, leading to deadlocks.
2.  **Add before Go**: Always call `Add()` in the parent goroutine *before* launching the goroutine. Calling `Add()` inside the goroutine is a race condition because the parent might reach `Wait()` before the scheduler even schedules the child.

---

## 🔍 Deep Dive 2: `sync.Mutex` and Normal vs. Starvation Modes

A `sync.Mutex` is a mutual exclusion lock. Go's Mutex implementation is highly optimized and operates in two modes:

### 1. Normal Mode
*   Goroutines waiting for the lock are lined up in a FIFO queue.
*   However, a newly arrived goroutine (already running on a CPU core) has an advantage: it does not have to park and resume. It can compete directly for the lock against the popped goroutine from the front of the queue.
*   **Result**: High throughput, but can starve older waiting goroutines.

### 2. Starvation Mode
*   If a goroutine fails to acquire the lock for more than **1 millisecond**, the Mutex transitions into **Starvation Mode**.
*   In this mode, ownership of the lock is handed over directly from the unlocking goroutine to the first waiter at the front of the queue.
*   New arriving goroutines do not attempt to acquire the lock and do not spin; they place themselves at the end of the wait queue immediately.
*   **Result**: Prevents starvation by bounding latency. Transitions back to normal mode once the wait queue is empty or a waiter acquires the lock in less than 1ms.

---

## 🔍 Deep Dive 3: `sync.RWMutex` (Reader-Writer Mutex)

An `RWMutex` allows multiple read locks simultaneously, but only one write lock at a time:
*   `RLock()` / `RUnlock()`: Shared lock. Multiple readers can hold it.
*   `Lock()` / `Unlock()`: Exclusive lock. Only one writer can hold it. Blocks all readers and writers.

**Rule of Thumb**: Use `sync.RWMutex` only if you have a **high ratio of reads to writes** (e.g., 90% reads, 10% writes). If reads and writes are roughly equal, a standard `sync.Mutex` is faster because it has simpler internal bookkeeping.

---

## 🔍 Deep Dive 4: `sync.Once` & Double-Checked Locking

`sync.Once` guarantees that a function is executed exactly once across the lifecycle of the program, even if called concurrently from thousands of goroutines:
```go
var once sync.Once

func setup() {
    fmt.Println("System initialized")
}

// Running this in 100 concurrent goroutines only prints once!
go once.Do(setup)
```
### Under the Hood
It uses an atomic variable and a Mutex to implement double-checked locking:
```go
if atomic.LoadUint32(&o.done) == 0 {
    o.doSlow(f)
}
```
This check is extremely fast because it avoids locking if the function has already executed.

---

## 🔍 Deep Dive 5: Data Race Detection (`-race`)

A **Data Race** occurs when two or more goroutines access the same memory location concurrently, at least one access is a write, and there is no synchronization (like locks or channels) protecting it.

Go has a built-in race detector powered by ThreadSanitizer. You can run tests or compile binaries with the `-race` flag:
```bash
go test -race ./...
go run -race main.go
```
If a race is detected, the runtime prints a detailed report showing the stack traces of the conflicting read and write operations.

---

## ⚠️ Common Gotchas

1.  **Copying Locks**: Just like WaitGroups, Mutexes cannot be copied. Copying a struct containing a Mutex copies the lock state (whether locked or unlocked), which leads to immediate deadlocks. Pass structs containing locks as pointers.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement concurrent counters protected by Mutexes and atomic operations. Test for races:
```bash
go test -race
```
Verify correctness.
