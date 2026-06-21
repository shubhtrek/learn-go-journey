# Step 3.2: Automated Unit Testing & Benchmarking 🧪

This step covers writing automated test suites, implementing idiomatic table-driven tests, profiling execution performance with benchmarks, and generating code coverage reports.

Official documentation:
*   [Go Command: Test packages](https://go.dev/doc/code#Testing)
*   [Go Package: testing](https://pkg.go.dev/testing)
*   [Go Blog: Table-driven tests](https://go.dev/wiki/TableDrivenTests)

---

## 🔍 Deep Dive 1: Standard Unit Testing Mechanics

Go has a built-in, first-class testing framework. You do not need third-party libraries (like Jest or JUnit) to write unit tests.

### Conventions
1.  **File Naming**: Test files must end with `_test.go` (e.g. `calc_test.go`).
2.  **Package Location**: Test files should live in the same package directory as the code they are testing.
3.  **Function Signature**: Test functions must begin with `Test` followed by an uppercase letter and accept a single parameter of type `*testing.T`:
    ```go
    func TestAdd(t *testing.T) {
        // ...
    }
    ```

### Assertion Actions
*   `t.Errorf(format, args...)`: Logs an error message and marks the test as **failed**, but continues executing the test.
*   `t.Fatalf(format, args...)`: Logs an error message and **halts** execution of the current test immediately.

---

## 🔍 Deep Dive 2: Table-Driven Tests (Idiomatic Go)

The idiomatic way to write tests in Go is using **Table-Driven Tests**. Instead of copy-pasting test functions for different inputs, you define a slice of anonymous structs (the "table") representing input parameters and expected outcomes, then loop over them:

```go
func TestFibonacci(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"base case zero", 0, 0},
        {"base case one", 1, 1},
        {"fibonacci five", 5, 5},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            result := Fibonacci(tc.input)
            if result != tc.expected {
                t.Errorf("Fibonacci(%d) = %d; want %d", tc.input, result, tc.expected)
            }
        })
    }
}
```
*   `t.Run(name, func)`: Runs a **subtest** under a distinct name. If one subtest fails, the other subtests continue running.

---

## 🔍 Deep Dive 3: Benchmarking & Performance Profiling

Go includes built-in benchmarking support to measure function execution speed and memory allocations.

### 1. Benchmark Signature
Benchmark functions must begin with `Benchmark` and accept a parameter of type `*testing.B`:
```go
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(20) // Function to measure
    }
}
```
*   `b.N`: The testing framework adjusts `b.N` dynamically until the loop runs for a long enough duration to calculate a stable average time.

### 2. Running Benchmarks
Run benchmarks in your terminal using the `-bench` flag:
```bash
go test -bench=.
```
Expected output:
```text
BenchmarkFibonacci-8    5000000        240 ns/op
```
This shows that the function completed in approximately 240 nanoseconds per operation on a machine running 8 CPU cores.

To check memory allocations, add the `-benchmem` flag:
```bash
go test -bench=. -benchmem
```

---

## 🔍 Deep Dive 4: Test Coverage Reports

Go provides tools to measure what percentage of your statement blocks are covered by tests.

### 1. View Coverage Percentage
```bash
go test -cover
```
### 2. Export Coverage Profiles
You can output a coverage profile and analyze it interactively in your web browser:
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```
This opens a browser tab showing your source code files with executed lines highlighted in green and untested lines in red.

---

## ⚠️ Common Gotchas

1.  **Test Cache**: By default, if package files haven't changed, Go caches test results and prints `(cached)` on subsequent runs. To bypass the test cache and force a complete run, use the `-count=1` flag:
    ```bash
    go test -count=1 ./...
    ```

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and implement tests for the challenge functions. Run:
```bash
go test -v
```
Verify all tests pass successfully.
