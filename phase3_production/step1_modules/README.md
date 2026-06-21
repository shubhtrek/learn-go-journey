# Step 3.1: Modules, Packages & Workspaces 📁

This step covers package architecture, access visibility, circular dependency prevention, and multi-module project workspace configurations.

Official documentation:
*   [Go Spec: Import Declarations](https://golang.org/ref/spec#Import_declarations)
*   [Go Command: multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
*   [Go Blog: Go Modules Reference](https://go.dev/ref/mod)

---

## 🔍 Deep Dive 1: Packages & Name Visibility

Go programs are constructed by linking together packages. Packages isolate namespace scopes and encourage encapsulation.

### The Capitalization Rule (Exporting)
In Go, access control is governed entirely by the first character of the identifier name:
*   **Exported Identifiers (Public)**: If an identifier (variable, constant, struct, interface, function) starts with a Unicode upper case letter (e.g. `NewClient`), it is exported and visible outside its package.
*   **Unexported Identifiers (Private)**: If it starts with a lower case letter or underscore (e.g. `clientConfig`), it is unexported and visible **only** within files of the same package.

There are no public, private, or protected keywords.

---

## 🔍 Deep Dive 2: Import Scopes and Side-Effect Imports

When you import a package:
```go
import "fmt"
```
The compiler looks for the package in your module dependencies or standard library.

### 1. Alias Imports
You can rename an import to avoid name collisions:
```go
import (
    randcrypto "crypto/rand"
    randmath   "math/rand"
)
```

### 2. Blank Imports (`_`)
If you import a package without calling its methods, the compiler throws an error. If you want to import it solely to run its `init()` initialization side-effects (e.g., registering database drivers), prefix it with an underscore:
```go
import _ "github.com/lib/pq" // Registers PostgreSQL driver in database/sql
```

### 3. Dot Imports (`.`)
If you import a package using a dot, the exported identifiers of that package are promoted directly into the file's lexical block (allowing you to call them without the package prefix). This is generally discouraged except in specific test configurations:
```go
import . "math"
// ...
x := Sin(0.5) // sin is imported directly from math!
```

---

## 🔍 Deep Dive 3: Circular Imports (Import Cycles)

Go **strictly forbids circular dependencies**. If Package `A` imports Package `B`, and Package `B` imports Package `A`, the compiler will refuse to build the program:
```text
import cycle not allowed
package A
    imports B
    imports A
```

### Why?
To keep compile times extremely fast and force clean, decoupled software architectures.

### Resolving Import Cycles
1.  **Extract Common Interfaces**: Move the common interfaces or data structs that both packages need to a third, independent package (e.g., `types` or `models`).
2.  **Use Interfaces**: Instead of Package `A` directly calling functions in Package `B`, have Package `A` define an interface representing the behavior it needs, and have Package `B` implement that interface and register itself.

---

## 🔍 Deep Dive 4: Multi-Module Workspaces (`go.work`)

When developing multiple interconnected local Go modules simultaneously (e.g. a shared library module and a web service module), managing dependencies via `go.mod` using `replace` directives can become cumbersome.

Go 1.18 introduced **Workspaces** via `go.work`:
*   A `go.work` file sits in the parent directory containing all module folders.
*   It overrides the module dependency resolution pipeline, letting the compiler resolve dependencies to your local modules directly.

### `go.work` Example:
```work
go 1.25.0

use (
    ./my-app
    ./my-shared-library
)
```
When you run `go build` inside `./my-app`, the compiler will automatically resolve imports referencing `my-shared-library` to your local folder instead of trying to download them from Git.

---

## ⚠️ Common Gotchas

1.  **Internal Packages**: Go supports an `internal/` directory convention. Any package inside a directory named `internal` can only be imported by packages that share the same parent directory root. This lets library maintainers hide implementation details from public consumption.

---

## 🎯 Practice Challenge
Open [practice.go](./practice.go) and verify package import behaviors. Run:
```bash
go run .
```
Verify compilation.
