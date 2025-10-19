# Go Basics - Learning Repository

This repository contains my journey learning Go programming fundamentals, from basic syntax to advanced concepts like concurrency, generics, and building REST APIs.

## 📚 Repository Structure

### Tutorial Section (`/tutorial/`)
Core Go programming concepts with practical examples:

#### **Basics** (`/basics/`)
- **File**: `main.go`
- **Topics**: Data types, variables, constants, type conversion
- **Key Concepts**: 
  - Integer overflow behavior
  - String manipulation and UTF-8 handling
  - Type casting and arithmetic operations
  - Variable declaration methods (`var`, `:=`)

#### **Arrays** (`/Arrays/`)
- **File**: `main.go`
- **Topics**: Array fundamentals and operations

#### **Slices** (`/Slices/`)
- **File**: `main.go`
- **Topics**: Dynamic arrays, slice operations, memory management
- **Key Concepts**: Slice as wrapper around arrays, automatic memory allocation

#### **Maps** (`/map_and_loop/`)
- **File**: `main.go`
- **Topics**: Hash maps, iteration patterns
- **Key Concepts**: Map behavior with missing keys

#### **Strings** (`/Strings/`)
- **File**: `main.go`
- **Topics**: String manipulation, immutability

#### **Pointers** (`/Pointers/`)
- **File**: `main.go`
- **Topics**: Memory addresses, pointer arithmetic

#### **Structs** (`/Struct/`)
- **File**: `main.go`
- **Topics**: Custom types, methods, interfaces
- **Key Concepts**:
  - Struct definition and initialization
  - Method receivers
  - Interface implementation
  - Anonymous structs

#### **Functions** (`/functions/`)
- **File**: `main.go`
- **Topics**: Function definitions, parameters, return values

#### **Goroutines** (`/Go_routines/`)
- **File**: `main.go`
- **Topics**: Concurrent programming, synchronization
- **Key Concepts**:
  - `go` keyword for concurrent execution
  - `sync.WaitGroup` for coordination
  - `sync.RWMutex` for thread-safe operations
  - Performance comparison (sequential vs concurrent)

#### **Channels** (`/channels/`)
- **Files**: 
  - `example/main.go` - Price monitoring simulation
  - `tut/main.go` - Channel fundamentals
- **Topics**: Inter-goroutine communication
- **Key Concepts**:
  - Channel creation and communication
  - `select` statement for non-blocking operations
  - Producer-consumer patterns

#### **Generics** (`/generics/`)
- **File**: `main.go`
- **Topics**: Type parameters, generic functions and structs
- **Key Concepts**:
  - Generic function with type constraints
  - Generic structs with type parameters
  - Type inference

### API Tutorial Section (`/api.tutorial/`)
Real-world REST API implementation:

#### **Project Structure**
```
api.tutorial/
├── cmd/api/main.go          # Application entry point
├── api/api.go              # API configuration
├── internal/
│   ├── handlers/           # HTTP request handlers
│   │   ├── api.go         # Route definitions
│   │   └── get_coin_balance.go
│   ├── middleware/         # Custom middleware
│   │   └── authorization.go
│   └── tools/             # Database utilities
│       ├── database.go
│       └── mockdb.go
├── go.mod                 # Dependencies
└── go.sum
```

#### **Dependencies**
- **Chi Router**: Lightweight HTTP router
- **Logrus**: Structured logging
- **Gorilla Schema**: Form data parsing

#### **Features**
- RESTful API endpoints
- Middleware implementation (authorization)
- Structured logging
- Database integration patterns
- Clean architecture separation

## 🚀 Getting Started

### Prerequisites
- Go 1.24.5 or later
- Basic understanding of programming concepts

### Running Tutorial Examples
```bash
# Navigate to tutorial directory
cd tutorial

# Run any example
go run basics/main.go
go run Go_routines/main.go
go run channels/example/main.go
```

### Running the API
```bash
# Navigate to API directory
cd api.tutorial

# Install dependencies
go mod tidy

# Run the API server
go run cmd/api/main.go
```

The API will start on `localhost:8000`

## 📖 Key Learning Concepts

### **Go Fundamentals**
- **Compiled Language**: Go compiles to binaries with minimal memory usage
- **Static Typing**: Strong type system with compile-time checking
- **Garbage Collection**: Automatic memory management
- **Built-in Concurrency**: Native support for concurrent programming

### **Data Types**
- **Primitives**: `bool`, `string`, `int`, `float`, `complex`
- **Integer Types**: `int8`, `int16`, `int32`, `int64`, `uint`, `uintptr`
- **Special Types**: `byte` (uint8), `rune` (int32)
- **Floating Point**: `float32`, `float64`
- **Complex Numbers**: `complex64`, `complex128`

### **Concurrency Concepts**
- **Goroutines**: Lightweight threads managed by Go runtime
- **Channels**: Thread-safe communication mechanism
- **Synchronization**: `sync.WaitGroup`, `sync.Mutex`, `sync.RWMutex`
- **Select Statement**: Non-blocking channel operations

### **Advanced Features**
- **Interfaces**: Implicit implementation, duck typing
- **Generics**: Type parameters for reusable code
- **Structs**: Custom types with methods
- **Pointers**: Memory management and references

## 📝 Learning Notes

The `Notes.txt` file contains comprehensive learning notes covering:
- Go language fundamentals
- Data type specifications
- Concurrency vs parallelism
- Memory management concepts
- Best practices and gotchas

## 🛠️ Development Tools

- **Go Modules**: Dependency management
- **Chi Router**: HTTP routing
- **Logrus**: Structured logging
- **Gorilla Schema**: Form parsing

## 📚 Learning Path

1. **Start with**: `basics/` - Understand Go syntax and data types
2. **Progress to**: `Struct/` - Learn object-oriented concepts
3. **Master**: `Go_routines/` and `channels/` - Concurrency fundamentals
4. **Explore**: `generics/` - Modern Go features
5. **Build**: `api.tutorial/` - Real-world application

## 🎯 Key Takeaways

- Go's simplicity and power for concurrent programming
- Strong typing system prevents many runtime errors
- Built-in concurrency makes it excellent for scalable applications
- Clean syntax and fast compilation
- Excellent for microservices and API development

---

*This repository represents a comprehensive learning journey through Go programming, from basic syntax to building production-ready APIs.*
