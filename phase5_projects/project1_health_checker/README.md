# Project 5.1: Concurrent URL Health Checker 🛠️

This project guides you through building a command-line tool that concurrently checks the HTTP health status of a list of target URLs. You will learn to optimize connection pooling, control worker counts, and handle OS interrupts gracefully.

Official documentation:
*   [Go Package: net/http](https://pkg.go.dev/net/http)
*   [Go Package: os/signal](https://pkg.go.dev/os/signal)

---

## 🔍 Deep Dive 1: HTTP Client Optimization

A common mistake in Go is using the default HTTP client (`http.DefaultClient`) or creating new clients for every request under high traffic. This leads to socket exhaustion (sockets stuck in `TIME_WAIT` state).

### Best Practices for HTTP Clients
1.  **Reuse Clients**: `http.Client` is safe for concurrent use by multiple goroutines. Create one client instance and share it.
2.  **Explicit Timeouts**: Never use a client without a timeout. By default, `http.DefaultClient` has no timeout, meaning a slow server can block your goroutine forever:
    ```go
    client := &http.Client{
        Timeout: 5 * time.Second, // Hard timeout for the entire request lifecycle
    }
    ```
3.  **Transport Tuning (Connection Pooling)**: Customize the underlying `http.Transport` to control connection limits:
    ```go
    transport := &http.Transport{
        MaxIdleConns:        100, // Maximum idle (keep-alive) connections across all hosts
        MaxIdleConnsPerHost: 10,  // Maximum idle connections to keep open per-host
        IdleConnTimeout:     90 * time.Second,
    }
    client := &http.Client{
        Transport: transport,
    }
    ```

---

## 🔍 Deep Dive 2: The Worker Pool Pattern

Launching one goroutine per URL is fine for 10 URLs. However, if you have 10,000 URLs, launching them all at once can overwhelm your local network, trigger OS file descriptor limits, or cause the target servers to rate-limit you.

To solve this, implement the **Worker Pool Pattern**:

```text
  [URLs Channel] ---> [ Worker 1 ]
                 ---> [ Worker 2 ] ---> [Results Channel]
                 ---> [ Worker 3 ]
```

1.  **Jobs Channel**: A buffered channel containing the URLs to check.
2.  **Workers**: A fixed number of goroutines (e.g., 5 workers) that read from the jobs channel, execute the HTTP request, and write results to a results channel.
3.  **Results Channel**: Receives the status of each URL.

This limits active concurrency to a stable, controllable count.

---

## 🔍 Deep Dive 3: Graceful Shutdown (OS Signal Handling)

A production-grade command-line tool must handle user interrupts (like pressing `Ctrl+C`) gracefully, closing open connections and saving progress before exiting.

Go supports catching system interrupts using the `os/signal` package:
*   Create a channel of type `os.Signal` with capacity 1.
*   Register it to receive specific signals (e.g., `os.Interrupt`, `syscall.SIGTERM`).
*   Wait on the signal channel using `select`.

```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

select {
case <-sigChan:
    fmt.Println("\nReceived shutdown signal. Cleaning up...")
    // Close channels, finish active jobs, exit cleanly
case <-allJobsDone:
    fmt.Println("All URLs checked successfully.")
}
```

---

## 🎯 Project Requirements

Your health checker tool must:
1.  Accept a list of URLs from a text file or command-line arguments.
2.  Use a worker pool to limit concurrent requests.
3.  Implement customized connection timeout and pooling.
4.  Handle `Ctrl+C` interrupt signals to print a summary of processed URLs before exiting.
5.  Output results in a formatted text layout indicating: URL, Status Code (or Error), and response time in milliseconds.
