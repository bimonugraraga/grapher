package main

import (
	"fmt"
	"log"
)

// GraphWithDeps extends Graph with proper dependency resolution
type GraphWithDeps struct {
	Graph
	dependencies map[string][]string // Tracks which nodes each node depends on
}

// NewGraphWithDeps creates a new graph with dependency support
func NewGraphWithDeps() *GraphWithDeps {
	return &GraphWithDeps{
		Graph:        *NewGraph(),
		dependencies: make(map[string][]string),
	}
}

// AddEdge adds a dependency between nodes (from -> to)
func (g *GraphWithDeps) AddEdge(from, to string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// Store dependencies in the correct direction: "to" depends on "from"
	g.dependencies[to] = append(g.dependencies[to], from)
	// Also store in original edges for visualization
	g.edges[from] = append(g.edges[from], to)
}

func (g *GraphWithDeps) AddEdgeV2(from []string, to string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	for _, s := range from {
		if g.nodes[s].Err != nil {
			log.Printf("node %s has error: %v", s, g.nodes[s].Err)
			return
		}
		g.dependencies[to] = append(g.dependencies[to], s)
		// Also store in original edges for visualization
		g.edges[s] = append(g.edges[s], to)
	}
}

type ResultsNode struct {
	NodeName string
	Result   interface{}
	Err      error
}

// ExecuteWithDependencies executes nodes respecting their dependencies
func (g *GraphWithDeps) ExecuteWithDependencies() (map[string]interface{}, error) {
	results := make(map[string]interface{})
	visited := make(map[string]bool)

	var execute func(string) error
	execute = func(nodeID string) error {
		if visited[nodeID] {
			return nil
		}
		visited[nodeID] = true

		// Execute dependencies first
		if deps, exists := g.dependencies[nodeID]; exists {
			for _, dep := range deps {
				if err := execute(dep); err != nil {
					return err
				}
			}
		}

		// Execute current node after dependencies are done
		result, err := g.ExecuteNode(nodeID)
		if err != nil {
			return fmt.Errorf("node %s execution failed: %v", nodeID, err)
		}

		nodeResult := ResultsNode{
			NodeName: nodeID,
			Result:   result,
			Err:      err,
		}
		results[nodeID] = nodeResult

		return nil
	}

	// Execute all nodes
	g.mu.RLock()
	nodeIDs := make([]string, 0, len(g.nodes))
	for nodeID := range g.nodes {
		nodeIDs = append(nodeIDs, nodeID)
	}
	g.mu.RUnlock()

	for _, nodeID := range nodeIDs {
		if !visited[nodeID] {
			if err := execute(nodeID); err != nil {
				return nil, err
			}
		}
	}

	return results, nil
}

// PrintDependencies prints the dependency information
func (g *GraphWithDeps) PrintDependencies() {
	g.mu.RLock()
	defer g.mu.RUnlock()

	fmt.Println("Dependency Structure:")
	for node, deps := range g.dependencies {
		fmt.Printf("%s depends on: %v\n", node, deps)
	}
	fmt.Println()
}
