package main

// This code implements a directed graph using an adjacency list representation.
// It allows you to add connections between nodes (represented by strings) and check if specific connections exist.
// The graph is directed because addEdge("A", "B") creates an edge from A to B, but not from B to A.
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}
