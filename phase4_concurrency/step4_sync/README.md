# Step 4.4: Sync Mutexes & WaitGroups 🔒

Use `sync.WaitGroup` to wait for goroutines to finish, and `sync.Mutex` to protect shared state.

## Key Concepts
1. **`sync.WaitGroup`**: `Add()`, `Done()`, and `Wait()` are used to coordinate termination.
2. **`sync.Mutex`**: `Lock()` and `Unlock()` ensure only one goroutine accesses a resource at a time.

## How to Run
```bash
go run main.go
```
