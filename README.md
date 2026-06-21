# 🐹 Go (Golang) Mastery Journal

Welcome to my personal Go systems programming archive! This repository documents my journey from setting up the Go SDK compiler to mastering concurrency models, memory layouts, network sockets, and open-source codebase contribution patterns.

---

## 💡 About This Journey
*   **The Goal**: Develop deep system-level intuition of the Go language runtime, write spec-compliant code, optimize memory and allocation footprints, and contribute to large-scale open-source projects (like `jaegertracing/jaeger`).
*   **The Method**: Systems-level conceptual study paired with hands-on practice.
*   **Pair Programming**: I am learning alongside my AI systems programming partner, **Antigravity**. I study compiler and runtime details, implement challenges in `practice.go` files, and get detailed code reviews to enforce idiomatic style.

---

## 📁 Repository Layout
Every topic folder contains:
*   `README.md`: Systems breakdown, under-the-hood compiler/runtime mechanics, specifications, and common gotchas.
*   `main.go`: A fully working demonstration showcasing the topic in action.
*   `practice.go`: My custom code implementation for that topic's coding challenge.

---

## 🗺️ Learning Roadmap & Progress

### ⛺ Phase 1: Zero to Hero (The Essentials)
- [ ] [Step 1.0: Environment, Installation & Toolchain](./phase1_basics/step0_setup) — Compiler pipeline, environment variables, modules, and GMP scheduler.
- [ ] [Step 1.1: Hello World, Syntax & Compilation](./phase1_basics/step1_hello) — Semicolon injection, package execution, and import initialization.
- [ ] [Step 1.2: Variables, Types & Zero Values](./phase1_basics/step2_variables) — Static type assertions, untyped constants, iota, and zero values.
- [ ] [Step 1.3: Control Flow (Spec-Compliant Loops)](./phase1_basics/step3_control) — Scoping, switches, range value copies, and Go 1.22+ loop variables.
- [ ] [Step 1.4: Functions & Memory Evaluation](./phase1_basics/step4_functions) — Named results, closures, and strict pass-by-value evaluation.
- [ ] [Step 1.5: Pointers & Escape Analysis](./phase1_basics/step5_pointers) — Stack vs. Heap allocation rules, composite literals, and compiler gcflags.

### 📦 Phase 2: Structuring Data & Logic (The Go Way)
- [ ] [Step 2.1: Slices vs. Arrays](./phase2_data_structures/step1_slices) — Slice headers, dynamic append allocation growth, and three-index slicing.
- [ ] [Step 2.2: Maps (Hash Map Internals)](./phase2_data_structures/step2_maps) — hmap buckets, lookup, non-addressability, and concurrency limits.
- [ ] [Step 2.3: Structs & Methods](./phase2_data_structures/step3_structs) — CPU word alignment padding, value/pointer receiver sets, and embedding composition.
- [ ] [Step 2.4: Interfaces](./phase2_data_structures/step4_interfaces) — iface/eface structures, type assertions, and the nil-interface trap.
- [ ] [Step 2.5: Error Handling](./phase2_data_structures/step5_errors) — errors.New, custom error structs, wrapping, and errors.Is/As.

### 🛠️ Phase 3: Writing Production Go
- [ ] [Step 3.1: Modules & Workspaces](./phase3_production/step1_modules) — Visibility capitalisation, side-effects import, cycles, and go.work.
- [ ] [Step 3.2: Automated Unit Testing](./phase3_production/step2_testing) — Table-driven subtests, testing.B benchmarks, and coverprofile.
- [ ] [Step 3.3: Defer, Panic & Recover](./phase3_production/step3_defer_panic) — Defer LIFO evaluation, panic propagation, and recover bounds.

### ⚡ Phase 4: Concurrency & The Superpowers (Advanced)
- [ ] [Step 4.1: Goroutines & The Scheduler](./phase4_concurrency/step1_goroutines) — OS Threads vs Goroutines, GMP work-stealing, and preemption.
- [ ] [Step 4.2: Channels (Thread-Safe Pipelines)](./phase4_concurrency/step2_channels) — hchan struct, wait queues, closed/nil states, and direction constraints.
- [ ] [Step 4.3: Select, Context & Timeouts](./phase4_concurrency/step3_select) - Random select polling, default clauses, time.After, and Context hierarchies.
- [ ] [Step 4.4: Sync Package Primitives](./phase4_concurrency/step4_sync) — WaitGroup counts, Mutex starvation modes, sync.Once, and race checks.

### 🚀 Phase 5: Projects & Open Source
- [ ] **Project 5.1**: [Concurrent URL Health Checker](./phase5_projects/project1_health_checker) — HTTP client transports, worker pools, and OS signals.
- [ ] **Project 5.2**: [High-Performance REST API](./phase5_projects/project2_todo_api) — net/http servers, Go 1.22+ routing parameters, middleware, and JSON codecs.

---

## 🏃 Quick Command Cheat-sheet

To run the program in any folder:
```bash
go run .
```

To run tests:
```bash
go test -v
```

To run test coverage check:
```bash
go test -cover
```

To check for data races:
```bash
go test -race ./...
```

To inspect compiler escape analysis:
```bash
go build -gcflags="-m"
```

---

*Made with 💖 by Shvbh*
