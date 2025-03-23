package dfs

import (
	"fmt"

	"github.com/martbul/4/graph"
)

func Dfs(node *graph.Node) {
	if node.Visited == true || node == nil {
		return
	}

	//visit node
	node.Visited = true
	fmt.Println(node.Value)

	for _, n := range node.Adjacent {
		Dfs(n)
	}

}
