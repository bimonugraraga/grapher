package main

import (
	"fmt"
)

func main() {
	graph := NewGraphWithDeps()
	graph.AddNode("calculation", func() (interface{}, error) {
		return 10 + 20, nil
	})
	results, _ := graph.ExecuteWithDependencies()

	fmt.Printf("\nAll results: %v\n\n", results)
	// graph.PrintDependencies()

}
