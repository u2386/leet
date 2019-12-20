package main

import (
	"fmt"
	"strings"
)

var graph map[string][]string

func generatePattern(word string) []string {
	patterns := []string{}
	for i := 0; i < len(word); i++ {
		sb := strings.Builder{}
		sb.WriteString(word[:i])
		sb.WriteString("*")
		sb.WriteString(word[i+1:])
		patterns = append(patterns, sb.String())
	}
	return patterns
}

func generateGraph(word string, graph map[string][]string) {
	for _, pattern := range generatePattern(word) {
		if _, ok := graph[pattern]; !ok {
			graph[pattern] = []string{}
		}
		graph[pattern] = append(graph[pattern], word)
	}
}

func isVisited(word string, path []string) bool {
	for _, p := range path {
		if word == p {
			return true
		}
	}
	return false
}

func bfs(begin, end string, graph map[string][]string) (paths [][]string) {
	q := [][]string{[]string{begin}}

	for len(q) > 0 {
		path := q[0]
		q = q[1:]

		word := path[len(path)-1]
		if word == end {
			paths = append(paths, path)
			continue
		}

		for _, v := range generatePattern(word) {
			for _, item := range graph[v] {
				if isVisited(item, path) {
					continue
				}
				newPath := append(path[:0:0], path...)
				newPath = append(newPath, item)
				q = append(q, newPath)
			}
		}
	}

	return paths
}

func main() {
	begin, end := "hit", "cog"
	graph = make(map[string][]string)
	for _, word := range []string{"hot", "dot", "dog", "lot", "log", "cog"} {
		generateGraph(word, graph)
	}
	generateGraph(begin, graph)
	fmt.Println(graph)

	fmt.Println(bfs(begin, end, graph))
}
