package main

import (
	"fmt"
	"sync"
)

// Node represents a graph node that executes a function
type Node struct {
	ID       string
	Function func() (interface{}, error)
	Result   interface{}
	Err      error
	Executed bool
}

// Graph represents a collection of nodes and their dependencies
type Graph struct {
	nodes map[string]*Node
	edges map[string][]string // adjacency list
	mu    sync.RWMutex
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
		edges: make(map[string][]string),
	}
}

// AddNode adds a new node with a function to the graph
func (g *Graph) AddNode(id string, fn func() (interface{}, error)) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.nodes[id] = &Node{
		ID:       id,
		Function: fn,
		Executed: false,
	}
}

// AddEdge adds a dependency between nodes (from -> to)
func (g *Graph) AddEdge(from, to string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.edges[from] = append(g.edges[from], to)
}

// ExecuteNode executes a single node and stores its result
func (g *Graph) ExecuteNode(nodeID string) (interface{}, error) {
	g.mu.RLock()
	node, exists := g.nodes[nodeID]
	g.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("node %s not found", nodeID)
	}

	if node.Function == nil {
		return nil, fmt.Errorf("node %s has no function", nodeID)
	}

	// Execute the function
	result, err := node.Function()
	if err != nil {
		node.Err = err
		return nil, err
	}
	node.Result = result
	node.Err = err
	node.Executed = true

	return result, nil
}

// ExecuteGraph executes all nodes in insertion order
func (g *Graph) ExecuteGraph() map[string]interface{} {
	results := make(map[string]interface{})

	g.mu.RLock()
	defer g.mu.RUnlock()

	for id := range g.nodes {
		result, err := g.ExecuteNode(id)
		if err != nil {
			results[id] = err
		} else {
			results[id] = result
		}
		fmt.Printf("Node %s executed: %v\n", id, results[id])
	}

	return results
}

// GetNodeResult returns the result of a specific node
func (g *Graph) GetNodeResult(nodeID string) (interface{}, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if node, exists := g.nodes[nodeID]; exists {
		return node.Result, node.Err
	}
	return nil, fmt.Errorf("node %s not found", nodeID)
}

// PrintGraph prints the graph structure
func (g *Graph) PrintGraph() {
	g.mu.RLock()
	defer g.mu.RUnlock()

	fmt.Println("Graph Structure:")
	for from, toNodes := range g.edges {
		fmt.Printf("%s -> %v\n", from, toNodes)
	}
	fmt.Println()
}
