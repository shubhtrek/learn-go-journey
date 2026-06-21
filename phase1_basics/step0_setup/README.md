# Step 1.0: Environment, Installation & Toolchain Setup ⚙️

Welcome to the true entry point of your Go journey. Before writing any code, it is critical to understand the environment you are executing in, how the Go compiler translates your text files into machine instructions, and how the Go runtime orchestrates memory and scheduling.

Official documentation:
*   [Go Installation Guide](https://go.dev/doc/install)
*   [How to Write Go Code](https://go.dev/doc/code)
*   [Go Modules Reference](https://go.dev/ref/mod)
*   [Command go Reference](https://go.dev/cmd/go/)

---

## 🔍 Deep Dive 1: The Go Compiler and Linker Pipeline

Go is a compiled, statically typed language. It does not run on an interpreter (like Python) or compile to intermediate bytecode for a virtual machine (like Java). It compiles directly to native OS/architectural machine code.

The compiler pipeline (`compile` tool under `go tool compile`) follows these stages:

1.  **Lexing & Parsing (Syntax Analysis)**: The compiler reads UTF-8 encoded `.go` source files, tokenizes them, and builds a concrete syntax tree.
2.  **AST Construction**: The concrete syntax tree is transformed into an **Abstract Syntax Tree (AST)**. In this stage, identifiers, types, and function calls are checked for structural validity.
3.  **Type Checking & Semantic Analysis**: The AST is evaluated to ensure type safety. Go's type-checker enforces rules such as checking that types match in variable assignments, verifying function call arguments, and detecting unused variables or imports.
4.  **Static Single Assignment (SSA) Code Generation**: The AST is converted into a SSA form, an intermediate representation where every variable is assigned exactly once. The compiler performs target-independent optimizations here, such as dead code elimination, loop unrolling, and inlining small functions.
5.  **Machine Code Generation**: The optimized SSA representation is converted into CPU-specific assembly instructions (e.g., AMD64, ARM64).
6.  **Linking (`go tool link`)**: The linker combines the compiled package object files and the embedded Go Runtime into a single, self-contained executable binary file.

---

## 🔍 Deep Dive 2: Environment Variables Demystified

The Go toolchain relies on environment variables to control compilation, dependencies, and execution target parameters. You can inspect your current environment settings by running `go env`.

*   **`GOROOT`**: The directory where the Go SDK itself is installed (e.g., `C:\Program Files\Go` on Windows or `/usr/local/go` on macOS/Linux). You should rarely modify this manually; it is automatically determined by the installation.
*   **`GOPATH`**: The root of the Go workspace. In the legacy GOPATH-mode (prior to Go Modules), all projects and third-party packages lived here. In modern Go, it acts as a local cache directory for downloading modules and storing globally installed binaries.
    *   `GOPATH/src`: (Legacy) Where source code resided.
    *   `GOPATH/pkg/mod`: Where downloaded third-party modules are cached.
    *   `GOPATH/bin`: Where compiled executable binaries installed via `go install` are placed.
*   **`GOBIN`**: The directory where `go install` writes compiled binaries. If unset, it defaults to `$GOPATH/bin` (or `%GOPATH%\bin` on Windows).
*   **`GOOS` & `GOARCH`**: Target Operating System (e.g., `windows`, `linux`, `darwin`) and Target Architecture (e.g., `amd64`, `arm64`, `386`). Go supports **cross-compilation** out-of-the-box. You can compile a Linux binary from Windows by setting:
    ```bash
    $env:GOOS="linux"; $env:GOARCH="amd64"; go build -o myapp
    ```
*   **`GOPROXY`**: The URL used to fetch dependencies. By default, it points to `https://proxy.golang.org`, a public module mirror operated by Google to ensure fast downloads and package availability.
*   **`GOPRIVATE`**: A comma-separated list of glob patterns (matching import paths) that bypasses the public proxy and checksum database. This is used for internal enterprise repositories (e.g., `github.com/my-org/*`).

---

## 🔍 Deep Dive 3: The Go Runtime

Every Go binary includes the **Go Runtime** statically linked inside it. The runtime starts executing before your `main.main` function runs. It is responsible for:

1.  **The GMP Scheduler**: Go implements an M:N scheduler to handle concurrency:
    *   **G (Goroutine)**: Represents the goroutine execution state, stack, and program counter.
    *   **M (Machine)**: Represents an OS thread managed by the kernel.
    *   **P (Processor)**: Represents a logical resource required to execute Go code. The number of Ps is typically set to `runtime.GOMAXPROCS` (defaults to the CPU core count).
    The scheduler multiplexes thousands of Goroutines (G) onto a limited number of OS threads (M) using logical processors (P). It implements work-stealing and cooperative/asynchronous preemption to ensure fair execution.
2.  **Memory Allocator**: Based on the TCMalloc design, the allocator manages memory via thread-local caches (`mcache`) to avoid global lock contention, intermediate caches (`mcentral`), and a heap manager (`mheap`).
3.  **Garbage Collector (GC)**: A concurrent, tri-color mark-and-sweep collector. It runs in the background concurrently with application threads (mutators) to minimize "Stop the World" pauses (which are typically less than 1 millisecond).

---

## 🛠️ Essential Go CLI Tooling

Here is the breakdown of the essential command line utilities provided by the `go` command:

*   **`go version`**: Prints the installed Go toolchain version.
*   **`go env`**: Prints Go environment variables.
*   **`go run <path>`**: Compiles and runs a temporary binary. Perfect for quick testing.
*   **`go build`**: Compiles the package in the current directory and generates an executable in the same directory.
*   **`go install`**: Compiles and places the resulting executable in `$GOBIN` (or `$GOPATH/bin`), making it available globally in your terminal.
*   **`go fmt`**: Automatically formats your code according to the official style guidelines (using tabs, space formatting, alignment).
*   **`go vet`**: Statically analyzes the code to find potential issues that compile successfully but might behave incorrectly (e.g., mismatched Printf verbs, unreachable code).
*   **`go doc <symbol>`**: Prints documentation for standard library elements or custom code right inside your terminal.
*   **`go clean`**: Cleans up object files and temporary build directories.
*   **`go mod init <module-name>`**: Initializes a new Go module, creating a `go.mod` file.
*   **`go mod tidy`**: Scans your code, adds missing dependencies to `go.mod`, and prunes unused ones.

---

## ⚠️ Common Setup Gotchas

1.  **Path Configurations**: If `$GOPATH/bin` (or `%USERPROFILE%\go\bin` on Windows) is not in your system `PATH` variable, CLI tools installed via `go install` will fail to run when called by name.
2.  **Mixed Go Modules and GOPATH**: Trying to run code outside a Go module without initializing `go.mod` in modern Go versions (Go 1.16+) will lead to errors resolving external packages. Always run `go mod init` inside your project root.

---

## 🎯 Practice Challenge

Open [practice.go](./practice.go) and implement a diagnostic command-line tool that:
1. Imports the standard library `runtime` and `os` packages.
2. Retrieves and prints:
   *   The installed Go runtime version (`runtime.Version()`).
   *   The host Operating System (`runtime.GOOS`).
   *   The host CPU architecture (`runtime.GOARCH`).
   *   The number of logical CPUs allocated to the program (`runtime.NumCPU()`).
   *   The environment variables `GOROOT` and `GOPATH` via `os.Getenv()`.
3. Compiles the program using `go build` and runs the resulting binary.
