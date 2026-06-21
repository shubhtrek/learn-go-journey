# Step 4.3: Select, Context & Timeouts (Multiplexing & Lifecycle Controls) 🚦

This step covers multiplexing multiple channel streams using the `select` statement, non-blocking operations, and managing background tasks using the standard `context` library.

Official documentation:
*   [Go Spec: Select statements](https://golang.org/ref/spec#Select_statements)
*   [Go Package: context](https://pkg.go.dev/context)
*   [Go Blog: Go Concurrency Patterns: Context](https://go.dev/blog/context)

---

## 🔍 Deep Dive 1: The `select` Statement & Random Polling Order

The `select` statement lets a goroutine wait on multiple communication operations. It blocks until one of its cases is ready to execute:

```go
select {
case msg1 := <-ch1:
    fmt.Println("Received:", msg1)
case ch2 <- msg2:
    fmt.Println("Sent:", msg2)
}
```

### Random Selection Algorithm
Unlike a standard switch statement that evaluates cases sequentially from top to bottom, `select` uses a **pseudo-random evaluation order**:
*   If multiple cases are ready simultaneously, the runtime selects one at random.
*   **Why?** To prevent starvation. If cases were evaluated sequentially, the first case would always dominate if it received frequent updates, starving the subsequent channels of CPU cycles.

---

## 🔍 Deep Dive 2: Non-blocking Select Operations

If no channels are ready and a `default` case is defined, the select statement executes the `default` block immediately without blocking:
```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
default:
    fmt.Println("No data available, moving on...")
}
```
This is useful for rate-limiting, polling loops, or non-blocking channel checks.

---

## 🔍 Deep Dive 3: Timeout Management (`time.After`)

To prevent operations from blocking indefinitely, you can multiplex them with a timer using `time.After()`, which returns a channel that sends the current time after a specified duration:
```go
select {
case res := <-webRequestChannel:
    fmt.Println("Data received:", res)
case <-time.After(2 * time.Second):
    fmt.Println("Timeout! Server did not respond in 2 seconds.")
}
```

---

## 🔍 Deep Dive 4: Context Propagation (`context.Context`)

The `context` package is the standard Go mechanism for propagating cancellation signals, deadlines, and request-scoped values across API boundaries and goroutine trees.

```text
       [ context.Background() ] (Root Context)
                  |
         [ Context with Cancel ]
         /                   \
[ Context with Timeout ]   [ Context with Value ]
```

### 1. Root Contexts
*   `context.Background()`: Returns an empty context. Usually used at the main entry point or top-level request handler.
*   `context.TODO()`: Used when you are unsure which context to use, or as a placeholder.

### 2. Context Cancellation (`WithCancel`)
Returns a copy of the parent context with a new `Done` channel, and a `cancel` function:
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // Always call cancel to release resources!

go func(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return // Stop execution when cancel is called
        default:
            // Do background work...
        }
    }
}(ctx)
```

### 3. Context Deadlines & Timeouts (`WithTimeout` & `WithDeadline`)
Automatically cancels the context when a duration expires or a specific timestamp is reached:
```go
ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
```
If the HTTP request takes longer than 500ms, the context cancels, and the client aborts the network connection.

---

## ⚠️ Common Gotchas

1.  **Leaking Contexts**: When you create a child context using `WithCancel`, `WithTimeout`, or `WithDeadline`, the runtime sets up a parent-child relationship in memory. If you do not call the returned `cancel()` function, the child context remains attached to the parent indefinitely, causing memory leaks. **Rule**: Always call `defer cancel()` immediately after creation.
2.  **Context for Values**: While `context.WithValue` exists, it should **never** be used to pass optional function arguments. It should only carry request-scoped values (such as trace IDs, security tokens, or authentication credentials).

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement select timeout handlers and context cancellations. Run:
```bash
go run .
```
Verify compiler output and timeout logs.
