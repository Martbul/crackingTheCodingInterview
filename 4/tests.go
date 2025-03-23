package main

import (
	"fmt"

	fourOne "github.com/martbul/4/4.1"
	//	bfs "github.com/martbul/4/BFS"
	//	dfs "github.com/martbul/4/DFS"
	"github.com/martbul/4/graph"
)

func main() {
	//root := graph.CreateGraph()

	//dfs.Dfs(root)

	//fmt.Println("-------------------------------")
	//bfs.BFS(root)

	startNode, endNode := graph.CreateDirectedGraph()

	res := fourOne.RouteBetweenNodes(startNode, endNode)
	fmt.Println(res)

}
