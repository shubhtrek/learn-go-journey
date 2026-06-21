# Step 4.2: Channels (Thread-Safe Pipelines & Runtime Structures) 📞

This step covers Go's channel implementation, the memory layout of the `hchan` runtime struct, buffered vs. unbuffered communication, and safe channel lifecycles.

Official documentation:
*   [Go Spec: Channel Types](https://golang.org/ref/spec#Channel_types)
*   [Go Spec: Send / Receive operators](https://golang.org/ref/spec#Send_statements)
*   [Go Blog: Share Memory By Communicating](https://go.dev/blog/codelab-share)

---

## 🔍 Deep Dive 1: CSP Paradigm & Memory Safety

Go concurrency is built on the **CSP (Communicating Sequential Processes)** model:
> "Do not communicate by sharing memory; instead, share memory by communicating."

Channels act as synchronous, thread-safe pipes that allow goroutines to pass typed data without explicit locks or race conditions.

---

## 🔍 Deep Dive 2: Channel Internals (`hchan` struct)

Under the hood, a channel is a pointer to a struct named `hchan` (defined in `runtime/chan.go`). It uses a lock internally, but wraps it in a clean, language-level API.

```text
hchan struct memory layout:
+-------------------+-------------------+-------------------+
|       qcount      |      dataqsiz     |       buf ptr     |
| (current len)     | (capacity buffer) | (circular buffer) |
+-------------------+-------------------+-------------------+
|      elemsize     |      closed       |     elemtype      |
|  (size of type)   | (boolean flag)    | (type metadata)   |
+-------------------+-------------------+-------------------+
|     sendx / recvx |      recvq        |       sendq       |
| (buffer indices)  | (waiting receivers| (waiting senders  |
|                   |  linked list)     |  linked list)     |
+-------------------+-------------------+-------------------+
```

### Key Components of `hchan`
1.  **Circular Ring Buffer (`buf`)**: An array representing the buffer for buffered channels.
2.  **Lock (`lock`)**: A mutex protecting all fields in `hchan` from concurrent access.
3.  **Wait Queues (`recvq` and `sendq`)**: Doubly-linked lists of waiting goroutines (represented by `sudog` structs) that are currently blocked trying to write to a full channel or read from an empty channel.

### Execution Walkthrough: Sending on a Full Channel
1.  Goroutine `G1` tries to send a value to channel `ch`.
2.  `G1` acquires the channel's internal lock.
3.  The runtime sees the buffer is full.
4.  `G1` allocates a `sudog` struct, packages its execution state, and appends it to the `sendq` queue.
5.  `G1` calls `gopark()`, yielding the thread to the GMP scheduler. `G1` is now in the **waiting** state.
6.  When another goroutine reads from `ch`, it pops `G1` from `sendq`, copies the data directly from `G1`'s stack to its own (bypassing the buffer entirely for speed!), and calls `goready(G1)` to put `G1` back on the scheduler run queue.

---

## 🔍 Deep Dive 3: Unbuffered vs. Buffered Channels

### 1. Unbuffered Channels (`make(chan T)`)
*   **Capacity** = 0.
*   **Synchronization**: Blocks until both the sender and receiver are ready. A send operation blocks until a corresponding receive operation begins, and vice versa. It acts as a guarantee of synchronization (rendezvous).

### 2. Buffered Channels (`make(chan T, capacity)`)
*   **Capacity** > 0.
*   **Synchronization**: Does not block the sender if the buffer has free slots (`qcount < dataqsiz`). It only blocks if the buffer is full (for senders) or empty (for receivers).

---

## 🔍 Deep Dive 4: Channel Operations State Matrix

Understanding what happens during operations on channels in different states is critical to prevent deadlocks and runtime panics.

| Channel State | Send (`ch <- v`) | Receive (`<-ch`) | Close (`close(ch)`) |
| :--- | :--- | :--- | :--- |
| **`nil`** (Uninitialized) | **Blocks forever** | **Blocks forever** | **Panics** (`panic: close of nil channel`) |
| **Open & Empty** | Succeeds (blocks if unbuffered) | Blocks | Succeeds (receivers get zero-value and `ok = false`) |
| **Open & Buffered (Partial)**| Succeeds immediately | Succeeds immediately | Succeeds (buffered items can still be read) |
| **Open & Full** | Blocks | Succeeds immediately | Succeeds |
| **Closed** | **Panics** (`panic: send on closed channel`) | Succeeds (returns zero-value immediately, `ok = false`) | **Panics** (`panic: close of closed channel`) |

---

## 🔍 Deep Dive 5: Directional Channels (API Design)

To restrict channel privileges and enforce API safety, you can specify channel direction in parameters:

*   **Bidirectional**: `chan T` (can send and receive).
*   **Send-Only**: `chan<- T` (can only write, cannot close or read).
*   **Receive-Only**: `<-chan T` (can only read, cannot write or close).

```go
// producer can only write to the channel
func producer(out chan<- int) {
    out <- 42
    close(out) // Allowed
}

// consumer can only read from the channel
func consumer(in <-chan int) {
    val := <-in
    // in <- 10 // ❌ Compile error: send to receive-only type
}
```

---

## ⚠️ Common Gotchas

1.  **Leaking Goroutines**: If a goroutine is blocked sending to a channel that has no receiver, or receiving from a channel that has no sender, and the channel is never closed, the goroutine remains in memory forever. This is a memory leak. Always design clear ownership and exit pipelines for goroutines.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement synchronous and buffered pipelines. Run:
```bash
go run .
```
Verify compilation and synchronization behavior.
