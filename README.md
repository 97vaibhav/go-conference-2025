# After go func(): Goroutines Through a Beginner’s Eye

> A beginner-friendly exploration of how Go schedules Goroutines and what really happens when you write `go func()`

## 🎤 Speaker Deck

[View the presentation slides here](https://speakerdeck.com/97vaibhav/after-go-func-goroutines-through-a-beginners-eye-6c9ccb98-e933-4d79-8474-b1715a0f9f2f)

## 📝 Description

Have you ever wondered how Goroutines actually run behind the scenes? Why your concurrent Go code works… or sometimes doesn't?

### Background
When I first started writing concurrent programs in Go, I was amazed by how simple it was to launch a Goroutine – just one keyword! But soon I found myself asking: **"What really happens after I write `go func()`?"**

This talk is a beginner-friendly dive into how Go schedules Goroutines using its lightweight runtime scheduler. We'll explore:
- How Goroutines are mapped onto OS threads
- What M:N scheduling really means
- How Go handles preemption
- Why your code may behave differently than expected under load

As someone who transitioned from traditional programming models to Go's concurrency model, I'll share my personal learning journey including moments of confusion and clarity and how understanding Go's scheduler helped me write better, more efficient code.

This session is especially helpful for those who have used Goroutines without really knowing what's under the hood and want to take that first step toward understanding concurrency the Go way.

### Expected Effect on Audience
Attendees (especially beginners) will leave with:
- ✅ A clear mental model of how Go schedules Goroutines
- ✅ An understanding of key concepts like G, M, P, and work-stealing
- ✅ Tips for writing more predictable and efficient concurrent Go code
- ✅ Curiosity to explore Go's runtime deeper with tools like GODEBUG and scheduler traces

This talk bridges the gap between **"I can write Goroutines"** and **"I understand how they run."**

## 🚀 Demo Code

This repository contains practical demonstrations of Go's scheduler behavior that accompany the presentation.

### Demo 1: Basic Goroutine Scheduling (`demo1/`)
A simple example showing how multiple Goroutines are scheduled and executed concurrently.

```go
// Run the demo
cd demo1
go run demo1.go
```

**What it demonstrates:**
- Basic Goroutine creation and synchronization
- How Goroutines execute concurrently
- The role of `sync.WaitGroup` in coordination

### Demo 2: Scheduler Tracing (`demo2/`)
An advanced example that generates a trace file to visualize Go's scheduler behavior with both CPU-bound and IO-bound Goroutines.

```go
// Run the demo
cd demo2
go run demo2.go

// Analyze the trace
go tool trace trace.out
```

**What it demonstrates:**
- CPU-bound vs IO-bound Goroutine behavior
- How `runtime.Gosched()` provides preemption points
- Goroutine parking and unparking during IO operations
- Multiple P (processor) scheduling with `GOMAXPROCS(4)`
- Scheduler trace visualization

## 🛠️ Prerequisites

- Go 1.24.4 or later
- Basic understanding of Go syntax
- Curiosity about concurrent programming!

## 📊 Understanding the Trace Output

After running Demo 2, you can explore the trace visualization in your browser:

1. The trace shows Goroutines (G), OS threads (M), and logical processors (P)
2. CPU-bound Goroutines appear as continuous blocks of execution
3. IO-bound Goroutines show as intermittent activity with sleep periods
4. You can observe work-stealing between different P's

## 🔍 Key Concepts Covered

### The GMP Model
- **G (Goroutine)**: Lightweight threads managed by Go runtime
- **M (Machine)**: OS threads that execute Goroutines
- **P (Processor)**: Logical processors that schedule Goroutines onto M's

### Scheduler Features
- **Work-stealing**: P's can steal work from other P's queues
- **Preemption**: How Go runtime interrupts long-running Goroutines
- **System calls**: How blocking operations are handled
- **Network poller**: Efficient handling of network I/O

## 🎯 Running the Demos

### Quick Start
```bash
# Clone the repository
git clone https://github.com/97vaibhav/go-conference-2025.git
cd go-conference-2025

# Run basic demo
cd demo1 && go run demo1.go

# Run trace demo and open visualization
cd ../demo2 && go run demo2.go && go tool trace trace.out
```

### Environment Variables for Exploration
Try running the demos with these environment variables to see different scheduler behaviors:

```bash
# Show scheduler decisions
GODEBUG=schedtrace=1000 go run demo2.go

# Disable work-stealing
GOMAXPROCS=1 go run demo2.go

# Force scheduler to be more aggressive
GOGC=10 go run demo2.go
```

## 📚 Further Reading

- [Go Runtime Scheduler Documentation](https://golang.org/src/runtime/proc.go)
- [Understanding Go's Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)
- [Go Memory Model](https://golang.org/ref/mem)

## 🤝 Contributing

Found an interesting scheduler behavior or have questions? Feel free to open an issue or submit a pull request!


---

**Happy Gophering! 🐹**

*If this talk helped you understand Goroutines better, consider sharing it with other Go developers who might benefit from demystifying the scheduler.*
