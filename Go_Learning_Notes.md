# Go (Golang) Professional Learning Notes & Interview Prep

A comprehensive roadmap for mastering Go, from language fundamentals to advanced systems architecture and performance optimization. These notes are designed for sequential learning, architectural reference, and high-level interview preparation.

---

## 1) Go Basics
Go is a statically typed, compiled language designed for simplicity, concurrency, and performance. Its fundamentals define the syntax and structural integrity of every program.

### 1.1 Hello World + Structure
A standard Go executable requires a `package main` declaration and a `main()` function as the entry point. The `fmt` package is utilized for formatted I/O.
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

### 1.2 Variables
Variables in Go are named storage with explicit or inferred types. Inside functions, the short declaration operator `:=` is preferred for conciseness, while `var` is used for package-level declarations or when zero-value initialization is required.
```go
package main

import "fmt"

func main() {
	var age int = 25     // Explicit type
	name := "Anya"        // Type inference
	var score float64     // Zero value: 0
	fmt.Println(age, name, score)
}
```
*Zero values: `0` for numbers, `""` for strings, `false` for booleans, and `nil` for pointers/interfaces/slices/maps.*

### 1.3 Data Types
Go provides a robust set of numeric, string, and boolean types. Numeric types include signed (`int`) and unsigned (`uint`) integers of various sizes, and floating-point numbers (`float32/64`). `byte` is an alias for `uint8`, and `rune` represents a Unicode code point (alias for `int32`). Type conversions must be explicit.
```go
var b byte = 65    // 'A'
var r rune = 'R'   // Unicode code point
```

### 1.4 Constants
Constants are immutable values resolved at compile time. They are ideal for fixed configurations and numeric constants that remain unchanged throughout the application lifecycle.
```go
const Pi = 3.14159
const (
	Open = 1
	Close = 2
)
```

---

## 2) Control Structures
Go simplifies control flow by offering a minimal yet powerful set of constructs: `if`, `for`, and `switch`.

### 2.1 If / Else
Go supports standard conditional branching and includes "short statements," allowing a variable to be initialized and scoped strictly within the `if` block.
```go
if n := len(s); n > 0 {
	fmt.Println(n)
}
```

### 2.2 Loops
Go utilizes a single looping construct: `for`. It can function as a traditional C-style loop, a `while` loop, or an infinite loop.
```go
for i := 0; i < 3; i++ { ... } // Standard
for condition { ... }          // While-style
for { ... }                    // Infinite
```

### 2.3 Switch
The `switch` statement in Go is more readable than long `if-else` chains. It automatically breaks after each case unless `fallthrough` is explicitly used. It can also be used without an expression to act as a cleaner `if-else` ladder.
```go
switch day {
case "Mon", "Tue": fmt.Println("Weekday")
default: fmt.Println("Other")
}
```

---

## 3) Functions
Functions are the primary building blocks of Go logic, supporting multiple return values, named returns, and closures.

### 3.1 Multiple & Named Returns
Go functions can return multiple values, typically used for returning a result alongside an `error`. Named returns allow for cleaner documentation and implicit return statements.
```go
func div(a, b int) (result int, err error) {
	if b == 0 { return 0, errors.New("zero") }
	result = a / b
	return
}
```

### 3.2 Variadic Functions & Closures
Variadic functions accept a variable number of arguments, while closures allow functions to capture and maintain state from their outer scope—useful for middleware and stateful handlers.
```go
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
```

---

## 4) Collection Types: Arrays, Slices, Maps

### 4.1 Arrays & Slices
Arrays are fixed-size sequences where the length is part of the type. Slices are dynamic, flexible views into arrays.
**Internals:** A slice is a 24-byte header containing a pointer to the backing array, a length (`len`), and a capacity (`cap`). 
**Growth:** Go uses a smooth growth formula (doubling for small slices, ~1.25x for larger ones) to manage memory efficiently.
**Nil vs Empty:** A `nil` slice has no backing array, whereas an empty slice `[]int{}` has a zero-size backing array. Prefer `nil` for API responses unless an empty array is explicitly required.

### 4.2 Maps
Maps are hash tables providing O(1) average lookup. 
**Internals:** Maps use buckets (8 entries each) and AES-based hashing. They perform incremental evacuation during growth to prevent latency spikes.
**Concurrency:** Maps are NOT thread-safe. Concurrent R/W operations cause an unrecoverable `panic`. Use `sync.RWMutex` or `sync.Map` for concurrent access.

---

## 5) Structs, Methods, and Memory
Structs group fields into custom types, and methods attach behavior.
**Memory Alignment:** Go aligns fields to word size. To minimize padding and optimize memory, order fields from largest to smallest.
**Receivers:** Use pointer receivers (`*T`) to mutate the struct or avoid copying large data; use value receivers (`T`) for immutability.

---

## 6) Interfaces
Interfaces define behavior through method sets and are implemented implicitly. They are the key to decoupling, mocking, and writing highly testable code in Go. 
**Tip:** Keep interfaces small (often 1-3 methods) to maximize composition.

---

## 7) Pointers
Pointers store memory addresses. They allow for efficient data passing and in-place updates. Go does not support pointer arithmetic, ensuring memory safety while maintaining performance.

---

## 8) Error Handling
Go treats errors as values. Functions return `(T, error)`, and the caller is expected to check the error explicitly. Use `fmt.Errorf("... %w", err)` to wrap errors and maintain the original context for debugging.

---

## 9) Concurrency Model
Go's concurrency is based on CSP (Communicating Sequential Processes).

### 9.1 Goroutines & GMP Scheduler
Goroutines are lightweight threads (starting at 2KB stack) managed by the Go runtime. 
**GMP Model:**
- **G (Goroutine):** Unit of execution.
- **M (Machine):** OS Thread.
- **P (Processor):** Resource required to run Gs (defaults to `GOMAXPROCS`).
The scheduler uses **Work Stealing** and **Hand-off** to ensure high CPU utilization and non-blocking execution.

### 9.2 Channels & Select
Channels are typed conduits for communication. 
- **Unbuffered:** Synchronous handshake.
- **Buffered:** Asynchronous until capacity is reached.
The `select` statement allows a goroutine to wait on multiple channel operations, often combined with `context` for timeouts and cancellation.

---

## 10) Context & Performance
The `context` package propagates cancellation and deadlines across API boundaries. 
**Optimization:**
- **Escape Analysis:** The compiler determines if variables stay on the stack or escape to the heap.
- **sync.Pool:** Reuse objects to reduce GC pressure.
- **GC:** Uses a Tri-color Mark-and-Sweep algorithm optimized for low STW (Stop The World) latency.

---

## 11) Testing & Design Patterns
**Testing:** Go encourages table-driven tests for comprehensive coverage and `fuzzing` for edge-case discovery.
**Design Patterns:**
- **Functional Options:** Clean API for constructors with many optional parameters.
- **Middleware:** Decorator pattern for HTTP handlers.

---

# Senior Interview Prep

## A) Technical Deep Dive
1. **Unbuffered vs Buffered Channels:** Unbuffered requires a synchronous handshake; buffered allows async sends until the buffer is full.
2. **Channel Closures:** Writing to a closed channel panics; reading from one returns the zero value and `ok == false`.
3. **Duck Typing:** Go interfaces are implicit. If a type satisfies the method set, it implements the interface.
4. **Race Detection:** Use `go test -race` to detect unsynchronized shared memory access.
5. **Deadlocks:** Occur when all goroutines are blocked. Debug using `pprof` or stack traces.

## B) Architectural Scenarios
1. **Rate Limiting:** Implement using a token bucket via channels and tickers.
2. **Worker Pools:** Use a fixed set of goroutines reading from a shared jobs channel to bound resource usage.
3. **Graceful Shutdown:** Use `context` and `os/signal` to ensure in-flight requests finish before the process exits.

---

# Best Practices
- **Check Errors:** Never ignore an error return.
- **Small Interfaces:** Favor small, focused interfaces.
- **Concurrency:** Don't communicate by sharing memory; share memory by communicating.
- **Memory:** Pre-allocate slices and maps with `make` when the size is known to avoid unnecessary reallocations.
