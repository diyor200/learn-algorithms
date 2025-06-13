package graph

import (
	"fmt"
	"testing"
)

func bfs(graph map[string][]string, name string) string {
	queue := []string{}
	visited := map[string]bool{}

	queue = append(queue, graph[name]...)

	for len(queue) > 0 {
		person := queue[0]
		queue = queue[1:]

		if !visited[person] {
			if isSeller(person) {
				return person
			} else {
				queue = append(queue, graph[person]...)
				visited[person] = true
			}
		}
	}

	return ""
}

func isSeller(name string) bool {
	return len(name) > 0 && name[len(name)-1] == 'm'
}

func TestFindMangoSeller(t *testing.T) {
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}
	res := bfs(graph, "you")
	if res != "" {
		fmt.Println(res)
	} else {
		fmt.Println("No mango seller found.")
	}
}

func bfsShortestPath(graph map[string][]string, start, target string) []string {
	var (
		queue   = [][]string{{start}}
		checked = map[string]bool{}
	)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]

		if checked[node] {
			continue
		}

		checked[node] = true

		if node == target {
			return path
		}

		for _, neighbor := range graph[node] {
			newPath := append([]string{}, path...)
			newPath = append(newPath, neighbor)
			queue = append(queue, newPath)
		}
	}

	return nil
}

func TestBFSShortestPath(t *testing.T) {
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}

	path := bfsShortestPath(graph, "you", "jonny")
	if path != nil {
		fmt.Println("Shortest path:", path)
	} else {
		fmt.Println("No path found.")
	}

}
