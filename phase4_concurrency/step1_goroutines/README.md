# Step 4.1: Goroutines & The Go Runtime Scheduler 🏃

This step covers how Go implements lightweight concurrency, differences between concurrency and parallelism, and the architecture of the Go GMP scheduler.

Official documentation:
*   [Go Doc: FAQ - Stack or Heap](https://go.dev/doc/faq#stack_or_heap)
*   [Go Spec: Go statements](https://golang.org/ref/spec#Go_statements)
*   [Go Runtime Scheduler Design](https://go.dev/src/runtime/proc.go)

---

## 🔍 Deep Dive 1: Concurrency vs. Parallelism

Many developers use these terms interchangeably, but they represent fundamentally different paradigms:

*   **Concurrency**: The composition of independently executing processes (dealing with **structure**). A system is concurrent if it can handle multiple tasks by interleaving their execution on a single core. It is about *structure*.
*   **Parallelism**: The simultaneous execution of multiple processes at the exact same instant (dealing with **execution**). It requires multi-core hardware. It is about *execution*.

> "Concurrency is about structure, parallelism is about execution." — Rob Pike

---

## 🔍 Deep Dive 2: Goroutines vs. OS Threads

A goroutine is a lightweight thread managed by the Go runtime, not the Operating System.

| Aspect | OS Thread | Goroutine |
| :--- | :--- | :--- |
| **Memory footprint** | Fixed size, typically **1MB** to **8MB**. | Dynamic size, starts at only **2KB**. |
| **Scheduling** | Preempted by the OS kernel (expensive context switches requiring CPU register save/restore). | Managed by Go Runtime (cheap context switches, entirely in user space). |
| **Creation Cost** | Slow (requires kernel call and thread allocation). | Fast (nanoseconds, simple memory allocation). |
| **Scaling** | Limited to thousands before system degradation. | Millions of concurrent goroutines can run on a single laptop. |

### Dynamic Stack Resizing
To start with such a small footprint, Go uses **contiguous stacks**. When a goroutine runs out of its 2KB stack space (e.g. deep recursion), the Go runtime allocates a new, contiguous memory block that is twice the size, copies the old stack contents to the new block, adjusts all pointers inside the stack, and frees the old stack.

---

## 🔍 Deep Dive 3: The GMP Scheduler Model

The Go scheduler multiplexes goroutines across OS threads using the **GMP Model**:

```text
  [ G1 ]  [ G2 ]  [ G3 ]  (Goroutines Queue)
            |
            v
          [ P ] (Logical Processor)
            |
            v
          [ M ] (OS Thread)
            |
            v
       [ CPU Core ]
```

*   **G (Goroutine)**: Represents the goroutine. It holds the execution stack, instruction pointer, and scheduling state.
*   **M (Machine / OS Thread)**: Represents a physical operating system thread. M executes the instruction code on the CPU core.
*   **P (Processor)**: Represents a logical context or resource required to execute Go code. The number of Ps matches `runtime.GOMAXPROCS` (defaults to the system's logical core count).

### Key Scheduler Algorithms

#### 1. Local and Global Run Queues
*   Each **P** maintains a local run queue of runnable goroutines.
*   There is also a single global run queue for goroutines that have been yielded or pre-allocated.

#### 2. Work-Stealing Algorithm
If a logical processor **P** runs out of goroutines in its local queue:
1.  It checks the global run queue.
2.  If the global queue is empty, it randomly selects another **P** and steals **half** of its local run queue to keep the CPU cores fully utilized.

#### 3. Asynchronous Preemption (Go 1.14+)
Prior to Go 1.14, the scheduler was cooperative: a goroutine could only be preempted if it made a function call (where the compiler injected stack-split checks). A tight loop without function calls (e.g. `for {}`) could block a thread forever.
In modern Go, the runtime uses OS signals to preempt goroutines asynchronously every 10ms, preventing starvation of other tasks.

---

## ⚠️ Common Gotchas

1.  **Main Goroutine Exit**: When the `main` goroutine exits, the entire program terminates immediately. Any background goroutines that are still executing are instantly killed without cleanup:
    ```go
    func main() {
        go func() {
            time.Sleep(1 * time.Second)
            fmt.Println("Done") // Will never print!
        }()
        // Program exits here, killing the background goroutine
    }
    ```
    **Fix**: Use synchronization primitives (like channels or `sync.WaitGroup`) to ensure the main goroutine waits for background processes to complete.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement concurrent goroutine execution. Run:
```bash
go run .
```
Verify order of execution and concurrency.
