package fourOne

import (
	"fmt"

	"github.com/martbul/4/graph"
)

func RouteBetweenNodes(startNode *graph.Node, endNode *graph.Node) bool {
	res := dfs(startNode, endNode)
	return res
}

func dfs(node *graph.Node, endNode *graph.Node) bool {
	if node == nil || node.Visited == true {
		return false
	}

	if node == endNode {
		return true
	}

	node.Visited = true
	fmt.Println(node.Value)

	for _, n := range node.Adjacent {
		if dfs(n, endNode) {
			return true
		}
	}

	return false
}
