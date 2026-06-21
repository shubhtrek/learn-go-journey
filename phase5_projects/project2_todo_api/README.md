# Project 5.2: High-Performance REST API (Vanilla Go) 🌐

This project guides you through building a high-performance REST API using Go's standard `net/http` library. You will learn to use Go 1.22+ routing patterns, design middleware pipelines, and handle JSON serialization securely.

Official documentation:
*   [Go Package: net/http](https://pkg.go.dev/net/http)
*   [Go Blog: Routing Enhancements in Go 1.22](https://go.dev/blog/routing-enhancements)

---

## 🔍 Deep Dive 1: Modern Routing (Go 1.22+)

Prior to Go 1.22, the standard library multiplexer (`http.ServeMux`) was very basic: it did not support path parameters (e.g., `/todos/{id}`) or matching HTTP methods directly (e.g., `GET` vs `POST`), forcing developers to use third-party libraries like Gorilla Mux or Chi.

In Go 1.22+, `http.ServeMux` was enhanced to support these features natively:

### 1. Method Matching
You can prepend the HTTP method to the route pattern:
```go
mux := http.NewServeMux()
mux.HandleFunc("GET /todos", listTodosHandler)
mux.HandleFunc("POST /todos", createTodoHandler)
```

### 2. Path Parameters
You can define path wildcards using curly braces, and retrieve them using `PathValue`:
```go
mux.HandleFunc("GET /todos/{id}", getTodoHandler)

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id") // Extracts parameter value
    fmt.Fprintf(w, "Fetching todo: %s", id)
}
```

---

## 🔍 Deep Dive 2: Writing Custom HTTP Middleware

Middleware is a design pattern used to execute code before or after the main request handler (e.g. for logging, authentication, recovery, CORS).

### The Middleware Signature
In Go, middleware is simply a function that accepts an `http.Handler` and returns an `http.Handler`:

```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Call the next handler in the chain
        next.ServeHTTP(w, r)
        
        log.Printf("Method: %s | Path: %s | Duration: %s", r.Method, r.URL.Path, time.Since(start))
    })
}
```

### Chaining Middleware
```go
// Apply logging middleware to all routes handled by the mux
loggedMux := LoggingMiddleware(mux)
http.ListenAndServe(":8080", loggedMux)
```

---

## 🔍 Deep Dive 3: Safe JSON Serialization & Deserialization

When handling JSON requests in APIs:

### 1. Safe Decoding
Avoid reading the entire body into memory before parsing. Instead, decode directly from the request stream using `json.NewDecoder`:
```go
var todo Todo
err := json.NewDecoder(r.Body).Decode(&todo)
if err != nil {
    http.Error(w, "invalid request body", http.StatusBadRequest)
    return
}
```

### 2. Safe Encoding
Encode data directly back to the response writer:
```go
w.Header().Set("Content-Type", "application/json")
err := json.NewEncoder(w).Encode(todo)
```

---

## 🎯 Project Requirements

Your REST API must:
1.  Manage a list of Todo items in memory (struct-based storage) using thread-safe synchronization (like `sync.RWMutex`).
2.  Support CRUD operations using Go 1.22+ routing:
    *   `GET /todos` (List all)
    *   `GET /todos/{id}` (Get specific)
    *   `POST /todos` (Create item)
    *   `PUT /todos/{id}` (Update item)
    *   `DELETE /todos/{id}` (Delete item)
3.  Implement middleware:
    *   **Logger Middleware**: Logs method, path, response status, and duration.
    *   **Recovery Middleware**: Recovers from handler panics and returns an HTTP 500 Internal Server Error, preventing the entire server from crashing.
4.  Include unit tests for handlers using the `net/http/httptest` package.
