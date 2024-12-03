package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {

	fmt.Println(hasEdge("name", "liukun"))

	addEdge("name", "ikaros")
	fmt.Println(hasEdge("name", "ikaros"))

}

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
