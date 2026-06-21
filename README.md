# 🐹 Go (Golang) Mastery Journal

Welcome to my personal Go development archive! This repository documents my end-to-end journey from writing my first line of Go code to mastering advanced concurrency patterns, cloud-native architecture, and preparing for major open-source contributions.

---

## 💡 About This Journey
*   **The Goal**: Develop deep system level intuition of Go, write clean production code, master goroutines, and confidently contribute to large-scale open-source projects (like `jaegertracing/jaeger`).
*   **The Method**: Guided conceptual learning paired with hands-on practice. 
*   **Pair Programming**: I am learning alongside my AI coding mentor, **Antigravity**. I study the concepts, implement the code in `practice.go` files, and get manual code reviews to ensure idiomatic style and zero regrets!

---

## 📁 Repository Layout
Every topic folder contains:
*   `README.md`: Modern concept breakdown, mental models, under-the-hood mechanics, and common gotchas.
*   `main.go`: A fully working demonstration file showcasing the topic in action.
*   `practice.go`: My custom code implementation for that topic's coding challenge.

---

## 🗺️ Learning Roadmap & Progress

### ⛺ Phase 1: Zero to Hero (The Essentials)
- [ ] [Step 1.1: Hello World & Tooling](./phase1_basics/step1_hello) — Modules, compilation, and package main.
- [ ] [Step 1.2: Variables, Types & Zero Values](./phase1_basics/step2_variables) — Strong static typing, memory, and zero allocations.
- [ ] [Step 1.3: Control Flow (The One Loop)](./phase1_basics/step3_control) — If statements, switches, and the single unified loop.
- [ ] [Step 1.4: Functions & Multiple Returns](./phase1_basics/step4_functions) — Named returns, stack vs heap, and copy-by-value.
- [ ] [Step 1.5: Pointers Demystified](./phase1_basics/step5_pointers) — Memory addresses (`&`, `*`) and pointers.

### 📦 Phase 2: Structuring Data & Logic (The Go Way)
- [ ] [Step 2.1: Slices vs Arrays](./phase2_data_structures/step1_slices) — Backing arrays, slice headers, and capacity doubling.
- [ ] [Step 2.2: Maps (Key-Value Pairs)](./phase2_data_structures/step2_maps) — Hash maps and the comma-ok existence check.
- [ ] [Step 2.3: Structs & Methods](./phase2_data_structures/step3_structs) — Objects without classes: value vs pointer receivers.
- [ ] [Step 2.4: Interfaces](./phase2_data_structures/step4_interfaces) — Implicit contracts and loose decoupling.
- [ ] [Step 2.5: Error Handling](./phase2_data_structures/step5_errors) — Errors as values, no exceptions, explicit validation.

### 🛠️ Phase 3: Writing Production Go
- [ ] [Step 3.1: Modules & Packages](./phase3_production/step1_modules) — Package scoping, capitals for visibility, and imports.
- [ ] [Step 3.2: Unit Testing](./phase3_production/step2_testing) — Testing conventions and table-driven testing.
- [ ] [Step 3.3: Defer, Panic, & Recover](./phase3_production/step3_defer_panic) — Resource cleanup pipelines and panic safety.

### ⚡ Phase 4: Concurrency & The Superpowers (Advanced)
- [ ] [Step 4.1: Goroutines](./phase4_concurrency/step1_goroutines) — Multiplexed light threads on the Go scheduler.
- [ ] [Step 4.2: Channels](./phase4_concurrency/step2_channels) — Thread-safe communication pipelines over shared state.
- [ ] [Step 4.3: Select, Context & Timeouts](./phase4_concurrency/step3_select) — Multiplexing channels and context timeouts.
- [ ] [Step 4.4: Sync Package](./phase4_concurrency/step4_sync) — WaitGroups and Mutexes for concurrent safety.

### 🚀 Phase 5: Projects & Open Source
- [ ] **Project 1**: [Concurrent URL Health Checker](./phase5_projects/project1_health_checker)
- [ ] **Project 2**: [High-Performance REST API](./phase5_projects/project2_todo_api)

---

## 🏃 Quick Command Cheat-sheet

To run the program in any folder:
```bash
go run .
```

To initialize a new go module:
```bash
go mod init <module-name>
```

---

*Made with 💖 by Shubham*
