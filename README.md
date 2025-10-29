# Grapher

A lightweight Go library for building and executing computational graphs with dependency resolution.

## Quick Start

```bash
go get github.com/bimonugraraga/grapher
```

```go
package main

import (
    "fmt"
    "github.com/bimonugraraga/grapher"
)

func main() {
    // Create a basic graph
    graph := grapher.NewGraph()
    
    graph.AddNode("add", func() (interface{}, error) {
        return 10 + 20, nil
    })
    
    graph.AddNode("multiply", func() (interface{}, error) {
        return 5 * 8, nil
    })
    
    results := graph.ExecuteGraph()
    fmt.Printf("Results: %+v\n", results)
}
```

## Core Concepts

### Basic Graph (`grapher_vanilla.go`)

**Node**: Represents a computational unit with ID, function, and execution state.

**Graph**: Manages nodes and edges with thread-safe operations.

#### Key Functions:
- `NewGraph()` - Create new graph instance
- `AddNode(id, func)` - Add node with computation function
- `AddEdge(from, to)` - Define execution dependency
- `ExecuteNode(id)` - Execute single node
- `ExecuteGraph()` - Execute all nodes sequentially
- `GetNodeResult(id)` - Retrieve node execution result
- `PrintGraph()` - Visualize graph structure

### Dependency Graph (`grapher_with_depth.go`)

**GraphWithDeps**: Extends basic graph with proper dependency resolution.

#### Enhanced Functions:
- `NewGraphWithDeps()` - Create dependency-aware graph
- `AddEdgeV2(from[], to)` - Multiple dependencies support
- `ExecuteWithDependencies()` - Execute respecting dependency order
- `PrintDependencies()` - Show dependency relationships

## Examples

### Basic Usage

```go
// Simple arithmetic operations
graph := grapher.NewGraph()

graph.AddNode("sum", func() (interface{}, error) {
    return 15 + 25, nil  // 40
})

graph.AddNode("product", func() (interface{}, error) {
    return 6 * 7, nil    // 42
})

// Execute and get results
results := graph.ExecuteGraph()
```

### Dependency Management

```go
// Data processing pipeline
graph := grapher.NewGraphWithDeps()

graph.AddNode("fetch", func() (interface{}, error) {
    return []string{"data1", "data2"}, nil
})

graph.AddNode("process", func() (interface{}, error) {
    return "processed_data", nil
})

graph.AddNode("save", func() (interface{}, error) {
    return "saved", nil
})

// Define execution order: fetch → process → save
graph.AddEdge("fetch", "process")
graph.AddEdge("process", "save")

// Execute with dependency resolution
results, err := graph.ExecuteWithDependencies()
```

### Error Handling

```go
graph.AddNode("risky", func() (interface{}, error) {
    if someCondition {
        return nil, fmt.Errorf("operation failed")
    }
    return "success", nil
})

result, err := graph.ExecuteNode("risky")
if err != nil {
    log.Printf("Execution failed: %v", err)
}
```

## Features

- ✅ **Thread Safe**: `sync.RWMutex` protected operations
- ✅ **Error Propagation**: Proper error handling across nodes
- ✅ **Dependency Resolution**: Smart execution ordering
- ✅ **Visualization**: Graph structure printing
- ✅ **Flexible**: Support any return type via `interface{}`

## API Reference

### Graph Methods

| Method | Description | Returns |
|--------|-------------|---------|
| `AddNode(id, func)` | Add computation node | - |
| `AddEdge(from, to)` | Define dependency | - |
| `ExecuteNode(id)` | Execute single node | `(interface{}, error)` |
| `ExecuteGraph()` | Execute all nodes | `map[string]interface{}` |
| `GetNodeResult(id)` | Get node result | `(interface{}, error)` |

### GraphWithDeps Methods

| Method | Description |
|--------|-------------|
| `AddEdgeV2(from[], to)` | Multiple dependencies |
| `ExecuteWithDependencies()` | Dependency-aware execution |
| `PrintDependencies()` | Show dependency graph |

## Testing

Run the test suite:

```bash
go test -v ./...
```

Tests cover:
- Basic graph operations
- Dependency resolution
- Error handling scenarios
- Edge cases

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.