package bfs

import (
	"fmt"

	"github.com/martbul/4/graph"
)

func BFS(node *graph.Node) {
	queue := []*graph.Node{node}
	node.Visited = true

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		node.Visited = true
		fmt.Println(node.Value)

		for _, n := range node.Adjacent {
			queue = append(queue, n)
		}

	}
}
